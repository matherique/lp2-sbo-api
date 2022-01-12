package book

import (
	"context"
	"log"

	"github.com/matherique/lp2-sbo-api/internal/dto"
	"github.com/matherique/lp2-sbo-api/internal/models"
	"github.com/matherique/lp2-sbo-api/internal/repository"
	"github.com/matherique/lp2-sbo-api/pkg/utils"
)

type Book interface {
	Create(context.Context, dto.ReqNewBook) (*models.Book, error)
}

type book struct {
	repo repository.BookRepo
	log  *log.Logger
}

func New(repo repository.BookRepo) *book {
	b := new(book)
	b.log = utils.NewLogger("Book")
	b.repo = repo

	return b
}
