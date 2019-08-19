package data

import "time"

// Thread is a data model representing a thread
type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAt time.Time
}
