package controllers

import (
	"encoding/json"
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
	writeResponse(w, 200, notes)
}

// Find Action that finds note by id
func (NotesController) Find(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
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
	writeResponse(w, 200, note)
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

// Update Action that updates a note
func (NotesController) Update(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
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

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	params := allowedParams{}
	err = json.Unmarshal(body, &params)
	note.Title = params.Title
	err = note.Save()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
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

	writeResponse(w, 200, map[string]bool{"success": true})
}

func writeResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}
