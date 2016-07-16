package main

import (
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/jvikstedt/jnotes-backend/controllers"
	"github.com/jvikstedt/jnotes-backend/db"
	"github.com/urfave/negroni"
)

func main() {
	db.Setup()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := httprouter.New()

	notesController := controllers.NotesController{}

	router.GET("/", notesController.Index)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":" + port)
}
