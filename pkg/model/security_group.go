package model

import "time"

type SecurityGroup struct {
	Name        string    `dynamodbav:"name" json:"name,omitempty"`
	CreatedAt   time.Time `dynamodbav:"createdAt" json:"createdAt"`
	LastUpdated time.Time `dynamodbav:"lastUpdated" json:"lastUpdated"`
	Members     []User    `dynamodbav:"members" json:"members,omitempty"`
}
