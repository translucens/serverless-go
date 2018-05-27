package slackwebhook

import (
	"github.com/ashwanthkumar/slack-go-webhook"
)

type Incomming struct {
	text, channel string
}

// PostToSlack sends the messege to the incoming webhook URL
func PostToSlack(webhookURL string, payload slack.Payload) []error {

	return slack.Send(webhookURL, "", payload)
}
