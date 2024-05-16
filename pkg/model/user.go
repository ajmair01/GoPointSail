package model

import (
	"time"
)

type User struct {
	FirstName      string          `dynamodbav:"firstName" json:"firstName,omitempty"`
	LastName       string          `dynamodbav:"lastName" json:"lastName,omitempty"`
	MiddleInitial  string          `dynamodbav:"middleInitial" json:"middleInitial,omitempty"`
	UserName       string          `dynamodbav:"userName" json:"userName,omitempty"`
	Email          string          `dynamodbav:"email" json:"email,omitempty"`
	Phone          string          `dynamodbav:"phone" json:"phone,omitempty"`
	Password       string          `dynamodbav:"password" json:"password,omitempty"`
	CreatedAt      time.Time       `dynamodbav:"createdAt" json:"createdAt"`
	LastUpdated    time.Time       `dynamodbav:"lastUpdated" json:"lastUpdated"`
	SecurityGroups []SecurityGroup `dynamodbav:"securityGroups" json:"securityGroups,omitempty"`
	OtpCode        string          `dynamodbav:"otpCode" json:"otpCode,omitempty"`
	OtpExpiration  time.Time       `dynamodbav:"otpExpiration" json:"otpExpiration"`
	OptIn          OptIn           `dynamodbav:"optIn" json:"optIn"`
}
