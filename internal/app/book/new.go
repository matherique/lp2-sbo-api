package book

import (
	"context"
	"net/http"

	"github.com/matherique/lp2-sbo-api/internal/dto"
	"github.com/matherique/lp2-sbo-api/pkg/errors"
)

func (b *book) New(ctx context.Context, data dto.ReqNewBook) (*dto.ResNewBook, error) {
	return nil, errors.NewHttp(http.StatusNotFound, "not implemented")
}
