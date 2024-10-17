package commands

import (
	"context"
	"fmt"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// Verifying server to allow image generation commands
func SetupPermissions(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) *interactionError {
	// Bot needs to check permissions
	if ctx.Value(log.Key).(log.Fields)["author_id"] != viper.GetString("master") {
		s.ChannelMessageSend(i.ChannelID, "**Setup Unsucessful**. Permission denied.")
		return &interactionError{&PermissionDenied{}, "Permission Denied"}
	}

	viper.Set(fmt.Sprintf("verified.%s", i.GuildID), true)
	s.ChannelMessageSend(i.ChannelID, "**Start Dreaming! Sweet Dreams!** `~help`")
	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6, // Whisper Flag
			Content: "**Setup Complete!**",
		},
	}

	if err := s.InteractionRespond(i.Interaction, resp); err != nil {
		return &interactionError{err: err, message: "Couldn't reply to interaction."}
	}
	return nil
}
