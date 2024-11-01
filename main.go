package main

import (
	"Dreamer/api"
	"Dreamer/commands"
	"Dreamer/config"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var production *bool

func main() {
	// Sets Flag to Debug Mode
	production = flag.Bool("p", false, "enables production with json logging")
	flag.Parse()
	if *production {
		log.InitJSONLogger(&log.Config{Output: os.Stdout})
	} else {
		log.InitSimpleLogger(&log.Config{Output: os.Stdout})
	}

	// Sets up Configurations for Viper
	config.InitConfig()

	// Creates Discord Bot Session
	s, err := discordgo.New("Bot " + viper.GetString("discord.token"))
	if err != nil {
		log.WithError(err)
		return
	}

	// Sets up Gin Router
	go api.InitAPI(s)

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info("Bot has registered handlers")
	})

	// Listens to Messages
	//s.Identify.Intents = discordgo.IntentsGuildMessages
	//s.AddHandler(messageHandler)

	// Register Slash and Component Commands
	commands.RegisterSlashCommands(s)

	// Connecting to Discord Server Gateway
	s.Open()
	log.Info("Bot is initialising")
	// Goroutines which handles image generation queue
	go commands.StableDiffusionScheduler()
	go commands.UpdateQueueStatus()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	log.Info("Cleanly exiting")
	s.Close()
}

/**
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Info("Received message: " + m.Content + " from: " + m.Author.Username)
	// If message is sent from the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "Hello" {
		s.ChannelMessageSend(m.ChannelID, "World!")
	}
}
**/
