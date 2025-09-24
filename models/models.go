package models

import "time"

type User struct {
	ID        int64
	City      string
	CreatedAt time.Time
}
