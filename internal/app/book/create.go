package book

import (
	"context"
	"net/http"

	"github.com/matherique/lp2-sbo-api/internal/dto"
	"github.com/matherique/lp2-sbo-api/internal/models"
	"github.com/matherique/lp2-sbo-api/pkg/errors"
)

func (b *book) Create(ctx context.Context, data dto.ReqNewBook) (*models.Book, error) {
	if err := data.Validate(); err != nil {
		b.log.Printf("invalid data: %+v", err)
		return nil, errors.NewHttp(http.StatusBadRequest, err.Error())
	}

	book := data.ToBook()

	return book, nil
}
