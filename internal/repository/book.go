package repository

import (
	"github.com/matherique/lp2-sbo-api/internal/models"
	"github.com/matherique/lp2-sbo-api/pkg/store"
)

type BookRepo interface {
	Create(book models.Book) (*models.Book, error)
}

type bookRepo struct {
	store store.Store
}

func NewBookRepo(store store.Store) BookRepo {
	br := new(bookRepo)

	br.store = store

	return br
}

func (br *bookRepo) Create(book models.Book) (*models.Book, error) {
	return nil, nil
}
