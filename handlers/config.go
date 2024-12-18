package handlers

import "github.com/bwmarrin/discordgo"

// HandlerConfig handles configs for intents and handlers
func HandlerConfig(s *discordgo.Session) {
	s.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions
	s.AddHandler(MessageHandler)
	s.AddHandler(ReactionHandler)
}
