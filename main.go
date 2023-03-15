package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("SLACK_TOKEN environment variable is not set")
	}

	api := slack.New(token)

	// Send a message to a channel
	channelID, timestamp, err := api.PostMessage("#general", slack.MsgOptionText("Hello, world!", false))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
	fmt.Printf("Message sent to channel %s at %s\n", channelID, timestamp)

	// Receive messages from a channel
	conversations, _, err := api.ConversationsList(&slack.ConversationsListParams{})
	if err != nil {
		log.Fatalf("Error getting conversations: %v", err)
	}

	for _, conversation := range conversations {
		if conversation.Name == "general" {
			history, err := api.GetConversationHistory(&slack.GetConversationHistoryParameters{
				ChannelID: conversation.ID,
			})
			if err != nil {
				log.Fatalf("Error getting conversation history: %v", err)
			}
			for _, message := range history.Messages {
				fmt.Printf("Received message: %s\n", message.Text)
			}
		}
	}
}
