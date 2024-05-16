package model

import (
	"time"
)

type User struct {
	FirstName      string          `dynamodbav:"firstName"`
	LastName       string          `dynamodbav:"lastName"`
	MiddleInitial  string          `dynamodbav:"middleInitial"`
	UserName       string          `dynamodbav:"userName"`
	Email          string          `dynamodbav:"email"`
	Phone          string          `dynamodbav:"phone"`
	Password       string          `dynamodbav:"password"`
	CreatedAt      time.Time       `dynamodbav:"createdAt"`
	LastUpdated    time.Time       `dynamodbav:"lastUpdated"`
	SecurityGroups []SecurityGroup `dynamodbav:"securityGroups"`
	OtpCode        string          `dynamodbav:"otpCode"`
	OtpExpiration  time.Time       `dynamodbav:"otpExpiration"`
	OptIn          OptIn           `dynamodbav:"optIn"`
}
