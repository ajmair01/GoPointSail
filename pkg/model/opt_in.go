package model

type OptIn struct {
	AllowEmail bool `dynamodbav:"allowEmail"`
	AllowSMS   bool `dynamodbav:"allowSMS"`
}
