package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/mail"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/translucens/serverless-go/awscommon/snsevent"
	"github.com/translucens/serverless-go/intermediate"
)

var snsPostTopic string

func init() {
	snsPostTopic = os.Getenv("SNS_PUB_TOPIC")
}

func mailHandler(events snsevent.SNSEvent) error {

	for _, v := range events.Records {

		rawmsg := v.Sns.Message
		msg := new(snsevent.MailMsg)

		log.Printf("%s", rawmsg)

		if err := json.Unmarshal([]byte(rawmsg), msg); err != nil {
			return err
		}

		b64dec, err := base64.StdEncoding.DecodeString(msg.Base64Content)
		if err != nil {
			return err
		}

		parsed, err := mail.ReadMessage(bytes.NewReader(b64dec))
		if err != nil {
			return err
		}

		out := intermediate.MailContent{
			From:    parsed.Header.Get("From"),
			To:      parsed.Header.Get("To"),
			Subject: parsed.Header.Get("Subject"),
		}

		date, err := parsed.Header.Date()
		if err != nil {
			return err
		}
		out.Date = date

		buf := new(bytes.Buffer)
		buf.ReadFrom(parsed.Body)
		out.Body = buf.String()

		msgID, err := intermediate.PublishIntermediateToSNS(snsPostTopic, intermediate.FromMail, out)
		if err != nil {
			return err
		}

		log.Printf("published to SNS: %s\n", *msgID)
	}

	return nil
}

func main() {
	lambda.Start(mailHandler)
}
