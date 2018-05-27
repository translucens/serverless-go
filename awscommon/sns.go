package awscommon

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var (
	snsRegion = os.Getenv("SNS_REGION")
)

// PutSNS puts subject and message to SNS
// returns sent message ID, error
func PutSNS(topic, subject, message string) (*string, error) {

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	var input sns.PublishInput
	input.SetSubject(subject) // In the official document, it must be ASCII characters
	input.SetMessage(message)
	input.SetTopicArn(topic)

	svc := sns.New(sess)
	out, err := svc.Publish(&input)
	if err != nil {
		return nil, err
	}

	return out.MessageId, nil
}
