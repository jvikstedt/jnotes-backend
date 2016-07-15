package models

import (
	"github.com/jvikstedt/jnotes-backend/database"
	"time"
)

// Note represents Model Note
type Note struct {
	ID        int
	Title     string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// GetAllNotes Fetches all notes from the database
func GetAllNotes() []Note {
	notes := []Note{}
	database.DB.Select(&notes, "SELECT * FROM notes")
	return notes
}
