package model

import "time"

type Notification struct {
	ID        string
	AlertID   string
	Message   string
	Severity  string
	Channels  string
	Timestamp time.Time
	Sent      bool
}
