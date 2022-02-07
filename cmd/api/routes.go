package main

import (
	api "github.com/matherique/lp2-sbo-api/cmd/api/book"
	"github.com/matherique/lp2-sbo-api/internal/app"
	server "github.com/matherique/lp2-sbo-api/pkg/http"
	"github.com/matherique/lp2-sbo-api/pkg/store"
)

func loadRoutes(srv *server.Server, store store.Store) {
	// srv.Add("/healthcheck", healthcheck)

	book := api.NewBookApi(app.NewBookService(store))
	srv.AddAll(book)
}
