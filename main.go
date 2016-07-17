package main

import (
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/jvikstedt/jnotes-backend/controllers"
	"github.com/jvikstedt/jnotes-backend/db"
	"github.com/urfave/negroni"
)

func init() {
	db.Setup("./db/dbconf.yml")
}

func main() {
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
	router.GET("/notes/:id", notesController.Find)
	router.POST("/notes", notesController.Create)
	router.PATCH("/notes/:id", notesController.Update)
	router.DELETE("/notes/:id", notesController.Destroy)

	return router
}
