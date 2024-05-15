package models

import (
	"time"
)

type User struct {
	firstName      string
	lastName       string
	middleInitial  string
	userName       string
	email          string
	phone          string
	password       string
	createdAt      time.Time
	lastUpdated    time.Time
	securityGroups []SecurityGroup
	otpCode        string
	otpExpiration  time.Time
	optIn          OptIn
}
