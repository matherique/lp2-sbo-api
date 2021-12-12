package book

import (
	"context"

	"github.com/matherique/lp2-sbo-api/internal/dto"
	"github.com/matherique/lp2-sbo-api/internal/repository"
)

type Book interface {
	New(context.Context, dto.ReqNewBook) (*dto.ResNewBook, error)
}

type book struct {
	repo repository.BookRepo
}

func New(repo repository.BookRepo) *book {
	b := new(book)
	b.repo = repo

	return b
}
