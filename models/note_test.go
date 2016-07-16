package models

import (
	"github.com/jvikstedt/jnotes-backend/db"
	"testing"
)

func TestGetAllNotes(t *testing.T) {
	db.Setup("../db/dbconf.yml")
	notes := GetAllNotes()
	if len(notes) > 0 {
		t.Error("Expected length 0, got: ", len(notes))
	}
}
