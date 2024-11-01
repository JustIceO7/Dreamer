package queue

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

// Structure of each item
type Command struct {
	Session     *discordgo.Session
	Interaction *discordgo.InteractionCreate
	Prompt      string
	Priority    float64
}

// PriorityQueue for image generation commands
type PriorityQueue struct {
	queue        []Command
	mu           sync.Mutex
	ImageNotify  chan struct{}
	UpdateNotify chan struct{}
}

// Return length of priority queue
func (pq *PriorityQueue) Len() int {
	return len(pq.queue)
}

// Enqueues the next command
func (pq *PriorityQueue) Enqueue(s *discordgo.Session, i *discordgo.InteractionCreate, priority float64, prompt string) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	command := Command{
		Session:     s,
		Interaction: i,
		Prompt:      prompt,
		Priority:    priority,
	}
	idx := 0
	for idx < len(pq.queue) && pq.queue[idx].Priority > command.Priority {
		idx++
	}
	pq.queue = append(pq.queue[:idx], append([]Command{command}, pq.queue[idx:]...)...)

	select {
	case pq.ImageNotify <- struct{}{}:
	default:
	}
	select {
	case pq.UpdateNotify <- struct{}{}:
	default:
	}
}

// Dequeues next command from priority queue
func (pq *PriorityQueue) Dequeue() (Command, bool) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if len(pq.queue) > 0 {
		command := pq.queue[len(pq.queue)-1]
		pq.queue = pq.queue[:len(pq.queue)-1]
		return command, true
	}

	return Command{}, false
}

// Notifies image channel
func (pq *PriorityQueue) NotifyImageWorker() {
	if len(pq.queue) > 0 {
		select {
		case pq.ImageNotify <- struct{}{}:
		default:
		}
	}
}

// Notifies update channel
func (pq *PriorityQueue) NotifyUpdateWorker() {
	if len(pq.queue) > 0 {
		select {
		case pq.UpdateNotify <- struct{}{}:
		default:
		}
	}
}

// Return front of queue
func (pq *PriorityQueue) Peek() *Command {
	if len(pq.queue) > 0 {
		return &pq.queue[len(pq.queue)-1]
	}
	return nil
}

// Creates and returns an instance of a PriorityQueue
func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{
		queue:        make([]Command, 0),
		ImageNotify:  make(chan struct{}, 1),
		UpdateNotify: make(chan struct{}, 1),
	}
	return pq
}

// Returns queue
func (pq *PriorityQueue) CurrentQueue() []Command {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return pq.queue
}
