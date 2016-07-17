package main

import (
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/jvikstedt/jnotes-backend/controllers"
	"github.com/jvikstedt/jnotes-backend/db"
	"github.com/urfave/negroni"
)

func main() {
	db.Setup("./db/dbconf.yml")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	n := negroni.Classic()
	n.UseHandler(setupRouter())
	n.Run(":" + port)
}

func setupRouter() *httprouter.Router {
	router := httprouter.New()

	notesController := controllers.NotesController{}
	router.GET("/notes", notesController.Index)
	router.POST("/notes", notesController.Create)

	return router
}
