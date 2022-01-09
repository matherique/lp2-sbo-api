package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/matherique/lp2-sbo-api/internal/app/book"
	"github.com/matherique/lp2-sbo-api/internal/dto"
	"github.com/matherique/lp2-sbo-api/pkg/errors"
	server "github.com/matherique/lp2-sbo-api/pkg/http"
	"github.com/matherique/lp2-sbo-api/pkg/utils"
)

type bookApi struct {
	log *log.Logger
	app book.Book
}

var urlpattern = map[string]*regexp.Regexp{
	"/": regexp.MustCompile(`^\/books[\/]*$`),
}

func NewBookApi(app book.Book) *bookApi {
	ba := new(bookApi)
	ba.log = utils.NewLogger("Book")
	ba.app = app

	return ba
}

func (b *bookApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.log.Printf("%s - %s\n", r.Method, r.URL)

	var err error

	switch true {
	case r.Method == "GET" && urlpattern["/"].MatchString(r.URL.Path):
		err = b.getAll(w, r)
	case r.Method == "POST" && urlpattern["/"].MatchString(r.URL.Path):
		err = b.create(w, r)
	default:
		err = errors.NewHttp(http.StatusNotFound, "not found")
	}

	if err != nil {
		server.HttpError(w, err)
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

	server.Created(w, resp)
	return nil
}
