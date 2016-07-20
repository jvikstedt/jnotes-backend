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
	err = db.DB.Get(n, "INSERT INTO notes (title, created_at, updated_at) VALUES($1, $2, $3) RETURNING id, created_at, updated_at", n.Title, now, now)
	return
}

// Update updates a existing note
func (n *Note) Update() (err error) {
	now := time.Now()
	err = db.DB.Get(n, "UPDATE notes SET title=$1,updated_at=$2 WHERE id=$3 RETURNING updated_at", n.Title, now, n.ID)
	return
}

// Destroy Removes Note from database
func (n *Note) Destroy() (err error) {
	_, err = db.DB.Exec("DELETE FROM notes WHERE id=$1", n.ID)
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
func GetAllNotes() (notes []Note, err error) {
	notes = []Note{}
	err = db.DB.Select(&notes, "SELECT * FROM notes")
	return
}

// GetNote Find note by id
func GetNote(id int) (note Note, err error) {
	err = db.DB.Get(&note, "SELECT * FROM notes WHERE id=$1", id)
	return
}
