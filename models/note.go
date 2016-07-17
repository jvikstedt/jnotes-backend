package models

import (
	"time"

	"github.com/jvikstedt/jnotes-backend/db"
)

// Note represents Model Note
type Note struct {
	ID        int
	Title     string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Save Saves Note to database
func (n *Note) Save() (note Note, err error) {
	now := time.Now()
	err = db.DB.Get(&note, "INSERT INTO notes (title, created_at, updated_at) VALUES($1, $2, $3) RETURNING id, title, created_at, updated_at", n.Title, now, now)
	return
}

// GetAllNotes Fetches all notes from the database
func GetAllNotes() []Note {
	notes := []Note{}
	db.DB.Select(&notes, "SELECT * FROM notes")
	return notes
}
