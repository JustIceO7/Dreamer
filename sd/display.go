package sd

import (
	"bytes"
	"fmt"
	"time"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
)

// Displays generated images on Discord
func displayImages(s *discordgo.Session, i *discordgo.InteractionCreate, allImages *Images, timeTaken time.Duration) error {
	// Preparing images to be sent
	files := make([]*discordgo.File, len(allImages.Images))

	for n, imageData := range allImages.Images {
		reader := bytes.NewReader(imageData)
		fileName := fmt.Sprintf("image%d.png", n+1)
		files[n] = &discordgo.File{
			Name:   fileName,
			Reader: reader,
		}
	}

	// Create a message to be sent with the text and the images
	finishedContent := i.Interaction.Member.User.Mention() + " **Generation complete!** Time taken: " + timeTaken.String()

	// Edit the initial interaction response to include the generated images
	_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &finishedContent,
		Files:   files,
	})

	if err != nil {
		log.WithError(err).Error("Failed to edit interaction response with images.")
		return err
	}

	return nil
}
