package repository

import "github.com/matherique/lp2-sbo-api/pkg/store"

type BookRepo interface {
	// New(dto.NewBook) (*model.Book, error)
}

type bookRepo struct {
	store store.Store
}

func NewBookRepo(store store.Store) BookRepo {
	br := new(bookRepo)

	br.store = store

	return br
}
