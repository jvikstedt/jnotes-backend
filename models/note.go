package models

import (
	"github.com/jvikstedt/jnotes-backend/database"
	"time"
)

// Note model
type Note struct {
	ID        int
	Title     string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// GetAllNotes Fetch all notes
func GetAllNotes() []Note {
	notes := []Note{}
	database.DB.Select(&notes, "SELECT * FROM notes")
	return notes
}
