package models

import "time"

// Note contains the structure of a note
type Note struct {
	ID        uint      `json:"id"`
	OwnerName string    `json:"owner_name"`
	Title     string    `json:"title"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
