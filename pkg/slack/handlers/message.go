package slackhandlers

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"lingoose-test/pkg/llm"
)

func HandleMessageEvent(api *slack.Client, botUserID string, ev *slackevents.MessageEvent) {
	
	fmt.Printf("Message received from user: %s in channel: %s, text: %s\n", ev.User, ev.Channel, ev.Text)

	if ev.User != botUserID {

		llmResponse := llm.GenerateLLMResponse(ev.Text)

		_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText(llmResponse, false))
		if err != nil {
			fmt.Printf("Failed posting message: %v\n", err)
		}
	}
}
