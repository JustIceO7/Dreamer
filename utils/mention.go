package utils

import (
	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
)

// MentionUser sends ping message to user within a channel
func MentionUser(s *discordgo.Session, channelID string, user *discordgo.User, text string) error {
	_, err := s.ChannelMessageSend(channelID, user.Mention()+text)
	if err != nil {
		log.WithError(err).Error("Failed to send ping message.")
		return err
	}
	return nil
}
