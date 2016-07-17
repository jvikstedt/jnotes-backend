package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	notes, _ := models.GetAllNotes()
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

// Destroy Action that destroys a note
func (NotesController) Destroy(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	note, err := models.GetNote(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	note.Destroy()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", "{\"success\":\"true\"}")
}
