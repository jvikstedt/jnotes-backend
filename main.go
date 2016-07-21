package main

import (
	"os"

	"github.com/jvikstedt/jnotes-backend/controllers"
	"github.com/jvikstedt/jnotes-backend/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func init() {
	db.Setup("./db/dbconf.yml")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e := echo.New()

	api := e.Group("/api/v1")
	api.GET("/notes", controllers.GetNotes)
	api.GET("/notes/:id", controllers.GetNote)
	api.PATCH("/notes/:id", controllers.UpdateNote)
	api.DELETE("/notes/:id", controllers.DeleteNote)
	api.POST("/notes", controllers.CreateNote)
	e.Run(standard.New(":" + port))
}
