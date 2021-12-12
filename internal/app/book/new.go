package book

import (
	"context"
	"net/http"

	"github.com/matherique/lp2-sbo-api/internal/dto"
	"github.com/matherique/lp2-sbo-api/internal/models"
	"github.com/matherique/lp2-sbo-api/pkg/errors"
)

func (b *book) New(ctx context.Context, data dto.ReqNewBook) (*models.Book, error) {
	if err := validate(data); err != nil {
		return nil, err
	}

	book := data.ToBook()

	return book, nil
}

func validate(data dto.ReqNewBook) error {
	switch true {
	case data.Name == "":
		return errors.NewHttp(http.StatusBadRequest, "missing name")
	case data.ShortDesc == "":
		return errors.NewHttp(http.StatusBadRequest, "missing short_desc")
	case data.PaperBack == 0:
		return errors.NewHttp(http.StatusBadRequest, "missing paperback")
	case data.PublisherId == 0:
		return errors.NewHttp(http.StatusBadRequest, "missing publisher_id")
	case data.AuthorId == 0:
		return errors.NewHttp(http.StatusBadRequest, "missing author_id")
	case data.CategoryId == 0:
		return errors.NewHttp(http.StatusBadRequest, "missing category_id")
	}

	return nil
}
