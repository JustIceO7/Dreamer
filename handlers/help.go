package handlers

import "github.com/bwmarrin/discordgo"

// HelpEmbedding creates the embedding for the help menu
func HelpEmbedding(s *discordgo.Session, m *discordgo.MessageCreate) {
	botAvatarURL := s.State.User.AvatarURL("64")
	helpEmbed := &discordgo.MessageEmbed{
		Title:       "Dreamer Help",
		Description: "My prefix for commands is `&`\nReact with ❌ (:\u200Bx\u200B:) to delete generated images!",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: botAvatarURL,
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "__Useful Commands__",
				Value:  "`/dream` - Generates an image given prompt.",
				Inline: false,
			},
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed)
}
