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
	e.Run(standard.New(":" + port))
}
