package models

import (
	"github.com/jvikstedt/jnotes-backend/db"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setup() {
	db.Setup("../db/dbconf.yml")
}

func teardown() {
	clearNotes()
	db.DB.Close()
}

func clearNotes() {
	db.DB.MustExec("DELETE FROM notes")
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func validNote() (n Note) {
	n.Title = "Golang"
	return
}

func TestSave(t *testing.T) {
	defer clearNotes()
	note := validNote()

	// New Record
	err := note.Save()
	assert.Nil(t, err)
	assert.NotZero(t, note.ID)
	assert.Equal(t, note.Title, validNote().Title)

	// Update
	note.Title = "Ruby"
	id, updatedAt := note.ID, note.UpdatedAt
	err = note.Save()
	assert.Nil(t, err)
	assert.Equal(t, note.ID, id)
	assert.Equal(t, note.Title, "Ruby")
	assert.True(t, note.UpdatedAt.After(updatedAt))
}

func TestGetAllNotes(t *testing.T) {
	notes := GetAllNotes()
	assert.Zero(t, len(notes))
}
