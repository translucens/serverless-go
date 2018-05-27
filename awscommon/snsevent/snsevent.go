package snsevent

import (
	"time"
)

// SNSEvent is the event object sent from SNS
type SNSEvent struct {
	Records []struct {
		EventVersion         string `json:"EventVersion"`
		EventSubscriptionArn string `json:"EventSubscriptionArn"`
		EventSource          string `json:"EventSource"`
		Sns                  struct {
			SignatureVersion  string    `json:"SignatureVersion"`
			Timestamp         time.Time `json:"Timestamp"`
			Signature         string    `json:"Signature"`
			SigningCertURL    string    `json:"SigningCertUrl"`
			MessageID         string    `json:"MessageId"`
			Message           string    `json:"Message"`
			MessageAttributes struct {
				Test struct {
					Type  string `json:"Type"`
					Value string `json:"Value"`
				} `json:"Test"`
				TestBinary struct {
					Type  string `json:"Type"`
					Value string `json:"Value"`
				} `json:"TestBinary"`
			} `json:"MessageAttributes"`
			Type           string `json:"Type"`
			UnsubscribeURL string `json:"UnsubscribeUrl"`
			TopicArn       string `json:"TopicArn"`
			Subject        string `json:"Subject"`
		} `json:"Sns"`
	} `json:"Records"`
}
