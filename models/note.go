package models

import "time"

// Note model
type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
