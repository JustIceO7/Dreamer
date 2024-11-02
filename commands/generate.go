package commands

import (
	"Dreamer/queue"
	"Dreamer/sd"
	"Dreamer/utils"
	"context"
	"strconv"
	"time"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var imageQueue = queue.NewPriorityQueue()

// Handles image generation requests
func GenerateImage(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) *interactionError {
	memberID := i.Interaction.Member.User.ID
	g, err := s.Guild(i.GuildID)
	if err != nil {
		return &interactionError{err: err, message: "Unable to query guild"}
	}

	if !utils.IsMaster(memberID) {
		// Checks if guild is registered
		if _, exists := viper.GetStringMap("verified")[g.ID]; !exists {
			return &interactionError{&PermissionDenied{}, "Permission Denied"}
		}
	}

	contentResponse := i.Interaction.Member.User.Mention() + " Image generation in progress! Please wait..."
	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: contentResponse,
		},
	}

	if err := s.InteractionRespond(i.Interaction, resp); err != nil {
		return &interactionError{err: err, message: "Couldn't reply to interaction."}
	}

	prompt := i.ApplicationCommandData().Options[0].StringValue()
	log.Info("Prompt: " + prompt)
	priority := float64(time.Now().UnixNano()) / 1e9
	if utils.IsMaster(memberID) {
		priority -= 31536000.0 // 1 year
	}
	imageQueue.Enqueue(s, i, priority, prompt)

	return nil
}

// Handles scheduling image generation commands to StableDiffusionInit
func StableDiffusionScheduler() {
	for {
		<-imageQueue.ImageNotify
		log.Info("Generating image...")
		if cmd := imageQueue.Peek(); cmd != nil {
			err := sd.StableDiffusionInit(cmd.Session, cmd.Interaction, cmd.Prompt)
			if err != nil {
				errorMessage := cmd.Interaction.Member.Mention() + " Whoopsies Stable Diffusion API **Crashed/Doesn't Exist** :P"
				_, err := cmd.Session.InteractionResponseEdit(cmd.Interaction.Interaction, &discordgo.WebhookEdit{
					Content: &errorMessage,
				})

				if err != nil {
					log.WithError(err).Error("Failed to edit interaction response.")
				}
				log.WithError(err).Error("Error occured trying to generate image.")
			}
		}
		imageQueue.Dequeue()
		imageQueue.NotifyImageWorker()
		imageQueue.NotifyUpdateWorker()
	}
}

// Handles updating discord response message for generation queue
func UpdateQueueStatus() {
	for {
		<-imageQueue.UpdateNotify
		currentQueue := imageQueue.CurrentQueue()
		pos := 1
		for i := len(currentQueue) - 1; i >= 0; i-- {
			cmd := currentQueue[i]
			if cmd.PositionInQueue != pos {
				changedMessage := cmd.Interaction.Member.Mention() + " Image generation in progress! Your position in queue **#" + strconv.Itoa(pos) + "**"
				_, err := cmd.Session.InteractionResponseEdit(cmd.Interaction.Interaction, &discordgo.WebhookEdit{
					Content: &changedMessage,
				})
				if err != nil {
					log.WithError(err).Error("Failed to edit interaction response.")
				}
				cmd.PositionInQueue = pos
			}
			pos++
		}
	}
}
