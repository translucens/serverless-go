package snsevent

import (
	"time"
)

// MailMsg is SNS messege body from Email via SES
type MailMsg struct {
	NotificationType string `json:"notificationType"`
	Mail             struct {
		Timestamp        time.Time `json:"timestamp"`
		Source           string    `json:"source"`
		MessageID        string    `json:"messageId"`
		Destination      []string  `json:"destination"`
		HeadersTruncated bool      `json:"headersTruncated"`
		Headers          []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"headers"`
		CommonHeaders struct {
			ReturnPath string   `json:"returnPath"`
			From       []string `json:"from"`
			Date       string   `json:"date"`
			To         []string `json:"to"`
			MessageID  string   `json:"messageId"`
			Subject    string   `json:"subject"`
		} `json:"commonHeaders"`
	} `json:"mail"`
	Receipt struct {
		Timestamp            time.Time `json:"timestamp"`
		ProcessingTimeMillis int       `json:"processingTimeMillis"`
		Recipients           []string  `json:"recipients"`
		SpamVerdict          struct {
			Status string `json:"status"`
		} `json:"spamVerdict"`
		VirusVerdict struct {
			Status string `json:"status"`
		} `json:"virusVerdict"`
		SpfVerdict struct {
			Status string `json:"status"`
		} `json:"spfVerdict"`
		DkimVerdict struct {
			Status string `json:"status"`
		} `json:"dkimVerdict"`
		DmarcVerdict struct {
			Status string `json:"status"`
		} `json:"dmarcVerdict"`
		Action struct {
			Type     string `json:"type"`
			TopicArn string `json:"topicArn"`
			Encoding string `json:"encoding"`
		} `json:"action"`
	} `json:"receipt"`
	Base64Content string `json:"content"`
}
