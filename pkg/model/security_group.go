package model

import "time"

type SecurityGroup struct {
	Name        string    `dynamodbav:"name"`
	CreatedAt   time.Time `dynamodbav:"createdAt"`
	LastUpdated time.Time `dynamodbav:"lastUpdated"`
	Members     []User    `dynamodbav:"members"`
}
