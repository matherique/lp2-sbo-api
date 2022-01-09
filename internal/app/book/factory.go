package book

import (
	"log"

	"github.com/matherique/lp2-sbo-api/internal/repository"
	"github.com/matherique/lp2-sbo-api/pkg/store"
)

func MakeBookApp(store store.Store) *book {
	repo := repository.NewBookRepo(store)
	app := New(repo, &log.Logger{})

	return app
}