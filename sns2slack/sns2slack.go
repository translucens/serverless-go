package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ashwanthkumar/slack-go-webhook"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/translucens/serverless-go/awscommon/snsevent"
	"github.com/translucens/serverless-go/intermediate"
	"github.com/translucens/serverless-go/slackwebhook"
)

var (
	slackWebhookURL = os.Getenv("SLACK_WEBHOOK_URL")
	postTemplate    = os.Getenv("POST_TEMPLATE")
)

// SlackPayloader is able to converted to Slack payload
type SlackPayloader interface {
	SlackPayload() slack.Payload
}

// Handler generates Slack posts from SNS events
func handler(rawevents snsevent.SNSEvent) error {

	contents, err := intermediate.UnmarshalSNSMessage(&rawevents)
	if err != nil {
		return err
	}

	for i, ev := range contents {

		slackmsg, ok := ev.(SlackPayloader)
		if !ok {
			return fmt.Errorf("this type is not supported to convert a Slack payload; %s", rawevents.Records[i].EventSource)
		}

		pl := slackmsg.SlackPayload()
		if errs := slackwebhook.PostToSlack(slackWebhookURL, pl); len(errs) != 0 {
			for _, err := range errs {
				log.Println(err)
			}
			return fmt.Errorf("error happened while processing %dth SNS event", i)
		}
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
