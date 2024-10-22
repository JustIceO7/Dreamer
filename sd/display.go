package sd

import (
	"bytes"
	"fmt"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
)

// Displays generated images on Discord
func displayImages(s *discordgo.Session, i *discordgo.InteractionCreate, allImages *Images) error {
	_, err := s.ChannelMessageSend(i.ChannelID, i.Interaction.Member.User.Mention()+" Image Generation Complete!")
	if err != nil {
		log.WithError(err).Error("Failed to send ping message.")
		return err
	}
	for n, imageData := range allImages.Image {
		reader := bytes.NewReader(imageData)
		fileName := fmt.Sprintf("image%d.png", n+1)
		_, err := s.ChannelFileSend(i.ChannelID, fileName, reader)
		if err != nil {
			log.WithError(err).Error("Failed to display image.")
			return err
		}
	}
	return nil
}
