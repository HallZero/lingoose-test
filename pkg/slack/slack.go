package slack

import (
	"fmt"
	"log"
	"os"

	slackhandlers "lingoose-test/pkg/slack/handlers"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func Slack() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}

	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot:  ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	authTest, err := api.AuthTest()
	if err != nil {
		log.Fatalf("Unable to authenticate bot: %s", err)
	}

	botUserID := authTest.UserID 

	go func() {
		for evt := range client.Events {
			switch evt.Type {

			case socketmode.EventTypeConnecting:
				fmt.Println("Connecting to Slack with Socket Mode...")

			case socketmode.EventTypeConnectionError:
				fmt.Println("Connection failed. Retrying later...")

			case socketmode.EventTypeConnected:
				fmt.Println("Connected to Slack with Socket Mode.")

			case socketmode.EventTypeEventsAPI:
				
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)

				if !ok {
					fmt.Printf("Ignored %+v\n", evt)
					continue
				}

				client.Ack(*evt.Request)

				if eventsAPIEvent.Type == slackevents.CallbackEvent {
					innerEvent := eventsAPIEvent.InnerEvent
					switch ev := innerEvent.Data.(type) {

					case *slackevents.MessageEvent:
						slackhandlers.HandleMessageEvent(api, botUserID, ev)

					default:
						fmt.Printf("Unhandled event: %v\n", ev)
					}
				}

			case socketmode.EventTypeSlashCommand:
				cmd, ok := evt.Data.(slack.SlashCommand)

				if !ok {
					fmt.Printf("Ignored %+v\n", evt)
					continue
				}

				slackhandlers.HandleSlashCommand(api, client, cmd, evt)
				
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
			}
		}
	}()

	client.Run()
}
