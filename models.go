package snsparser

type Entry struct {
	Records []Record `json:"Records"`
}

type Record struct {
	EventSource          string `json:"EventSource"`
	EventVersion         string `json:"EventVersion"`
	EventSubscriptionArn string `json:"EventSubscriptionArn"`
	SNS                  SNS    `json:"Sns"`
}

type SNS struct {
	Type             string `json:"Type"`
	MessageId        string `json:"MessageId"`
	TopicArn         string `json:"TopicArn"`
	Subject          string `json:"Subject"`
	Message          string `json:"Message"`
	Timestamp        string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
}
