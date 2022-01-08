package main

import (
	"log"
	"net/http"
	"os"

	"github.com/matherique/lp2-sbo-api/internal/app/book"
	"github.com/matherique/lp2-sbo-api/internal/repository"
	"github.com/matherique/lp2-sbo-api/pkg/store"
)

func loadRoutes(srv *http.ServeMux, store store.Store) *http.ServeMux {
	logger := log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lmsgprefix)

	repo := repository.NewBookRepo(store)
	app := book.New(repo, logger)

	srv.Handle("/books", newBookApi(logger, app))

	return srv
}
