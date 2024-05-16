package model

type OptIn struct {
	AllowEmail bool `dynamodbav:"allowEmail" json:"allowEmail,omitempty"`
	AllowSMS   bool `dynamodbav:"allowSMS" json:"allowSMS,omitempty"`
}
