package sd

import (
	"Dreamer/clock"
	"Dreamer/utils"

	"github.com/bwmarrin/discordgo"
)

// Handles image generation commands
func StableDiffusionInit(s *discordgo.Session, i *discordgo.InteractionCreate, prompt string) error {
	timer := clock.NewTimer()
	imageData, err := TextToImage(prompt)
	if err != nil {
		return err
	}

	// Decodes images and saves into allImages
	allImages := &Images{}
	for _, img := range imageData.Images {
		rawBinary, err := utils.Decrypt(img)
		if err != nil {
			return err
		}
		allImages.AppendImage(rawBinary)
	}
	timeTaken := timer.Now()

	// Displays Images on Discord
	err = displayImages(s, i, allImages, timeTaken)
	if err != nil {
		return err
	}

	return nil
}
