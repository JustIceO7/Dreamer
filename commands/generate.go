package commands

import (
	"Dreamer/sd"
	"context"

	"github.com/bwmarrin/discordgo"
)

// Handles image generation requests
func GenerateImage(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) *interactionError {
	/*
		g, err := s.Guild(i.GuildID)
		if err != nil {
			return &interactionError{err: err, message: "Unable to query guild"}
		}

			// Checks if guild is registered
			if value, exists := viper.GetStringMap("verified")[g.ID]; exists {
				verifiedStatus := value.(bool)
				if !verifiedStatus {
					return &interactionError{&PermissionDenied{}, "Permission Denied"}
				}
			} else {
				return &interactionError{&PermissionDenied{}, "Permission Denied"}
			}
	*/
	prompt := i.ApplicationCommandData().Options[0].StringValue()

	err := sd.StableDiffusionInit(s, i, prompt)
	if err != nil {
		return &interactionError{err: err, message: "Error occured trying to generate image."}
	}

	return nil
}
