package dto

import "github.com/matherique/lp2-sbo-api/internal/models"

type ReqNewBook struct {
	Name        string `json:"name"`
	AuthorId    int    `json:"author_id"`
	ShortDesc   string `json:"short_desc"`
	Longdesc    string `json:"long_desc"`
	CategoryId  int    `json:"category_id"`
	PaperBack   int    `json:"paperback"`
	PublisherId int    `json:"publisher_id"`
}

func (rnw ReqNewBook) ToBook() *models.Book {
	return &models.Book{
		Name:        rnw.Name,
		AuthorId:    rnw.AuthorId,
		ShortDesc:   rnw.ShortDesc,
		Longdesc:    rnw.Longdesc,
		CategoryId:  rnw.CategoryId,
		PaperBack:   rnw.PaperBack,
		PublisherId: rnw.PublisherId,
	}
}
