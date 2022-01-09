package main

import (
	"net/http"

	api "github.com/matherique/lp2-sbo-api/cmd/api/book"
	"github.com/matherique/lp2-sbo-api/internal/app/book"
	"github.com/matherique/lp2-sbo-api/pkg/store"
)

func loadRoutes(srv *http.ServeMux, store store.Store) *http.ServeMux {
	srv.HandleFunc("/healthcheck", healthcheck)

	srv.Handle("/books", api.NewBookApi(book.MakeBookApp(store)))

	return srv
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
