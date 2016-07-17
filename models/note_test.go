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
	err := note.Save()
	assert.Nil(t, err)
	assert.NotZero(t, note.ID)
	assert.Equal(t, note.Title, validNote().Title)
}

func TestGetAllNotes(t *testing.T) {
	notes := GetAllNotes()
	if len(notes) > 0 {
		t.Error("Expected length 0, got: ", len(notes))
	}
}
