package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/jvikstedt/jnotes-backend/models"
	"net/http"
)

// NotesController controller
type NotesController struct{}

// Index Action
func (NotesController) Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	notes := models.GetAllNotes()
	notesJSON, _ := json.Marshal(notes)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", notesJSON)
}
