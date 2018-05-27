package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	slack "github.com/ashwanthkumar/slack-go-webhook"
	"github.com/translucens/serverless-go/awscommon"
	"github.com/translucens/serverless-go/awscommon/snsevent"
)

const (
	// Mail type indicates its messege is from SES
	FromMail = "mail"
)

type Obj struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

type Mail struct {
	ContentType string      `json:"contentType"`
	Content     MailContent `json:"content"`
}

type MailContent struct {
	From    string    `json:"from"`
	To      string    `json:"to"`
	Subject string    `json:"subject"`
	Body    string    `json:"body"`
	Date    time.Time `json:"date"`
}

func (mc *MailContent) SlackPayload() slack.Payload {
	return slack.Payload{
		Username: mc.From,
		Text:     fmt.Sprintf("To: %s\nSubject: %s\n%s", mc.To, mc.Subject, mc.Body),
	}
}

// PublishIntermediateToSNS publish content to SNS
func PublishIntermediateToSNS(topic, contenttype string, content interface{}) (*string, error) {

	var jsoncontent string
	switch contenttype {
	case FromMail:
		mail, ok := content.(MailContent)
		if !ok {
			return nil, fmt.Errorf("contenttype and the content does not match; contenttype: %s", contenttype)
		}
		jsonb, err := json.Marshal(mail)
		if err != nil {
			return nil, err
		}
		jsoncontent = fmt.Sprintf("%s", jsonb)
	default:
		return nil, fmt.Errorf("unsupported contenttype: %s", contenttype)
	}

	return awscommon.PutSNS(topic, contenttype, jsoncontent)
}

// UnmarshalSNSMessage unmarshals the message in SNS event JSON
func UnmarshalSNSMessage(evs *snsevent.SNSEvent) ([]interface{}, error) {

	dsts := make([]interface{}, len(evs.Records))

	for i, ev := range evs.Records {

		var dst interface{}

		switch ev.Sns.Subject {
		case FromMail:

			rawmsg := ev.Sns.Message
			dst = new(MailContent)

			log.Println(rawmsg)

			json.Unmarshal([]byte(rawmsg), dst)

			mailmsg, ok := dst.(*MailContent)
			if !ok {
				return nil, fmt.Errorf("content type and the content does not match; content type: %s", ev.Sns.Subject)
			}

			dsts[i] = mailmsg

		default:
			return nil, fmt.Errorf("unsupported content type: %s", ev.Sns.Subject)
		}

	}
	return dsts, nil
}
