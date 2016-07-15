package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "test")
	})

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":" + port)
}
