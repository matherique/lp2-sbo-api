package main

import (
	"log"
	"net/http"
	"os"
)

func loadRoutes(srv *http.ServeMux) *http.ServeMux {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	srv.Handle("/book", newBookApi(logger))

	return srv
}
