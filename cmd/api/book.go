package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/matherique/lp2-sbo-api/internal/app/book"
	"github.com/matherique/lp2-sbo-api/internal/dto"
	"github.com/matherique/lp2-sbo-api/pkg/errors"
)

type bookApi struct {
	log *log.Logger
	app book.Book
}

var urlpattern = map[string]*regexp.Regexp{
	"/": regexp.MustCompile(`^\/books[\/]*$`),
}

func newBookApi(logger *log.Logger) *bookApi {
	ba := new(bookApi)
	ba.log = logger

	return ba
}

func (b *bookApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.log.Printf("[BOOK] %s - %s\n", r.Method, r.URL.Path)

	var err error

	switch true {
	case r.Method == "GET" && urlpattern["/"].MatchString(r.URL.Path):
		err = b.getAll(w, r)
	default:
		err = errors.NewHttp(http.StatusNotFound, "not found")
	}

	if err != nil {
		httpError(w, err)
		return
	}
}

func (b *bookApi) getAll(w http.ResponseWriter, r *http.Request) error {
	return errors.NewHttp(http.StatusNotFound, "not found")
}

func (b *bookApi) create(w http.ResponseWriter, r *http.Request) error {
	var data dto.ReqNewBook

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return errors.NewHttp(http.StatusBadRequest, err.Error())
	}

	resp, err := b.app.New(r.Context(), data)

	if err != nil {
		return err
	}

	created(w, resp)
	return nil
}
