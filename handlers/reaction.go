package handlers

import (
	"strings"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
)

// ReactionHandler handles reactions for discord messages
func ReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.Emoji.Name == "❌" {
		msg, err := s.ChannelMessage(r.ChannelID, r.MessageID)
		if err != nil {
			log.WithError(err).Error("Could not retrieve message.")
			return
		}
		if msg.Author.ID == s.State.User.ID && hasImageInMessage(msg) {
			err := s.ChannelMessageDelete(r.ChannelID, r.MessageID)
			if err != nil {
				log.WithError(err).Error("Could not delete message.")
			}
		}
	}
}

// Check if the attachment is an image
func isImageAttachment(attachment *discordgo.MessageAttachment) bool {
	return strings.HasSuffix(strings.ToLower(attachment.Filename), ".png")
}

// Check if the message has an image attachment or embed
func hasImageInMessage(msg *discordgo.Message) bool {

	// Check attachments for images
	for _, attachment := range msg.Attachments {
		if isImageAttachment(attachment) {
			return true
		}
	}

	return false
}
