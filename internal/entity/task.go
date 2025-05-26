package entity

import "time"

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	UserID      string    `json:"user_id"`
}
