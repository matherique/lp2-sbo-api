package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/matherique/lp2-sbo-api/cmd/api/book"
	"github.com/matherique/lp2-sbo-api/internal/app/book"
	"github.com/matherique/lp2-sbo-api/pkg/store"
)

func loadRoutes(srv *http.ServeMux, store store.Store) *http.ServeMux {
	logger := log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lmsgprefix)

	srv.HandleFunc("/healthcheck", healthcheck)
	srv.Handle("/books", api.NewBookApi(logger, book.MakeBookApp(store)))

	return srv
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
