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
func (n *Note) Save() (err error) {
	now := time.Now()
	if n.IsNew() {
		err = db.DB.Get(n, "INSERT INTO notes (title, created_at, updated_at) VALUES($1, $2, $3) RETURNING id, title, created_at, updated_at", n.Title, now, now)
	} else {
		err = db.DB.Get(n, "UPDATE notes SET title=$1,updated_at=$2 RETURNING title, updated_at", n.Title, now)
	}
	return
}

// IsNew returns true if record has not been saved to database otherwise false
func (n *Note) IsNew() bool {
	if n.ID == 0 {
		return true
	}
	return false
}

// GetAllNotes Fetches all notes from the database
func GetAllNotes() []Note {
	notes := []Note{}
	db.DB.Select(&notes, "SELECT * FROM notes")
	return notes
}
