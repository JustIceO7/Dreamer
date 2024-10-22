package commands

import (
	"Dreamer/sd"
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// Handles image generation requests
func GenerateImage(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) *interactionError {

	g, err := s.Guild(i.GuildID)
	if err != nil {
		return &interactionError{err: err, message: "Unable to query guild"}
	}

	if i.Interaction.Member.User.ID != viper.GetString("master") {
		// Checks if guild is registered
		if _, exists := viper.GetStringMap("verified")[g.ID]; !exists {
			return &interactionError{&PermissionDenied{}, "Permission Denied"}
		}
	}

	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6, // Whisper Flag
			Content: "**Image Generation in Process! Please Wait!**",
		},
	}

	if err := s.InteractionRespond(i.Interaction, resp); err != nil {
		return &interactionError{err: err, message: "Couldn't reply to interaction."}
	}

	prompt := i.ApplicationCommandData().Options[0].StringValue()
	fmt.Println(prompt)
	err = sd.StableDiffusionInit(s, i, prompt)
	if err != nil {
		return &interactionError{err: err, message: "Error occured trying to generate image."}
	}

	return nil
}
