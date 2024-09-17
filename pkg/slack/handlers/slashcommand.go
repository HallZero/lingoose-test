package slackhandlers

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)


func HandleSlashCommand(api *slack.Client, client *socketmode.Client, cmd slack.SlashCommand, evt socketmode.Event) {
	client.Debugf("Slash command received: %+v", cmd)

	_, _, err := api.PostMessage(
		cmd.ChannelID,
		slack.MsgOptionText("Hey! I'm the FAQ bot for CS! How can I help you today?", false),
	)

	if err != nil {
		fmt.Printf("Failed posting message: %v\n", err)
	}

	client.Ack(*evt.Request)
}
