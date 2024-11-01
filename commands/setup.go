package commands

import (
	"Dreamer/utils"
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// Verifying server to allow image generation commands
func SetupPermissions(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) *interactionError {
	// Bot needs to check permissions
	if !utils.IsMaster(i.Member.User.ID) {
		s.ChannelMessageSend(i.ChannelID, "Setup Unsuccessful. **Permission denied**.")
		return &interactionError{&PermissionDenied{}, "Permission Denied"}
	}
	content := "Already Setup!"
	if _, exists := viper.GetStringMap("verified")[i.GuildID]; !exists {
		viper.Set("verified."+i.GuildID, true)
		content = "Setup Complete! **Start Dreaming!** `&help`"
	}

	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	}

	if err := s.InteractionRespond(i.Interaction, resp); err != nil {
		return &interactionError{err: err, message: "Couldn't reply to interaction."}
	}
	return nil
}
