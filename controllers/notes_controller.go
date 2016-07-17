package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/jvikstedt/jnotes-backend/models"
)

type allowedParams struct {
	Title string
}

// NotesController All actions will be bound to this
type NotesController struct{}

// Index Action that returns all notes
func (NotesController) Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	notes := models.GetAllNotes()
	notesJSON, _ := json.Marshal(notes)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", notesJSON)
}

// Create Action that creates a note
func (NotesController) Create(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	params := allowedParams{}
	err = json.Unmarshal(body, &params)
	note := models.Note{Title: params.Title}
	note.Save()
}
