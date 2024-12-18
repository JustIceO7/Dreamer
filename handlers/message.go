package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// MessageHandler handles message commands
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// log.Info("Received message: " + m.Content + " from: " + m.Author.Username)
	// If message is sent from the bot
	if m.Author.ID == s.State.User.ID {
		return
	}
	prefix := viper.GetString("prefix")

	// Checking for presence of prefix
	if m.Content[0] == prefix[0] {
		messageContent := m.Content
		spaceIndex := len(messageContent)
		for i := 0; i < len(messageContent); i++ {
			if (messageContent[i]) == ' ' {
				spaceIndex = i
				break
			}
		}
		firstWord := messageContent[0:spaceIndex]
		switch firstWord {
		case prefix:
			s.ChannelMessageSend(m.ChannelID, "type `&help` to open help menu.") // invalid prefix command
		case prefix + "help":
			HelpEmbedding(s, m)
		}
	}

}
