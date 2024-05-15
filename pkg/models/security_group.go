package models

import "time"

type SecurityGroup struct {
	name        string
	createdAt   time.Time
	lastUpdated time.Time
	members     []User
}
