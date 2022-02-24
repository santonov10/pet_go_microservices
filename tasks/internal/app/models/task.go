package models

import "time"

type Task struct {
	ID          string
	UserID      string
	TimeCreated time.Time
	TimeUpdated time.Time
	Header      string
	Description string
}
