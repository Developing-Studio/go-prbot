package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type PRHandler struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
}

func (self PRHandler) handleListCommand(args []string) {
	self.session.ChannelMessageSend(self.message.ChannelID, "inside list command function")
}
func (self PRHandler) handlePRCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	self.session = s
	self.message = m

	//remove /pr and get args
	var args []string = strings.Split(strings.TrimPrefix(m.Content, "/pr "), " ")
	command, args := args[0], args[1:]
	if command == "list" {
		self.handleListCommand(args)
	} else {
		self.printPRHelp()
	}
}

func (self PRHandler) printPRHelp() {
	var helpMsg string = `
Invalid command for /pr.
  options:
    help
    list
    add
    delete
    describe
	`
	self.session.ChannelMessageSend(self.message.ChannelID, helpMsg)
}
