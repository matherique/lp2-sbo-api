package book

import (
	"context"
	"log"

	"github.com/matherique/lp2-sbo-api/internal/dto"
	"github.com/matherique/lp2-sbo-api/internal/models"
	"github.com/matherique/lp2-sbo-api/internal/repository"
)

type Book interface {
	New(context.Context, dto.ReqNewBook) (*models.Book, error)
}

type book struct {
	repo repository.BookRepo
	log  *log.Logger
}

func New(repo repository.BookRepo, log *log.Logger) *book {
	b := new(book)
	b.log = log
	b.repo = repo

	return b
}
