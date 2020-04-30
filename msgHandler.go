package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If we received a PR command
	if strings.HasPrefix(m.Content, "/pr ") {
		prHandler := PRHandler{
			message: m,
			session: s}

		prHandler.handlePRCommands(s, m)
	}
}
