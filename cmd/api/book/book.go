package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/matherique/lp2-sbo-api/internal/app"
	"github.com/matherique/lp2-sbo-api/internal/dto"
	server "github.com/matherique/lp2-sbo-api/pkg/http"
	"github.com/matherique/lp2-sbo-api/pkg/utils"
)

type bookApi struct {
	log *log.Logger
	a   *app.Book
}

var urlpattern = map[string]*regexp.Regexp{
	"/":     regexp.MustCompile(`^\/books[\/]*$`),
	"/{id}": regexp.MustCompile(`^\/books\/(\d{1,})$`),
}

func NewBookApi(app *app.Book) *bookApi {
	ba := new(bookApi)
	ba.log = utils.NewLogger("Book")
	ba.a = app

	return ba
}

func (b *bookApi) Routes() map[string]http.Handler {
	return map[string]http.Handler{
		`GET \/books[\/]*$`:      http.HandlerFunc(b.getAll),
		`POST \/books[\/]*$`:     http.HandlerFunc(b.create),
		`PUT \/books[\/]*$`:      http.HandlerFunc(b.update),
		`GET \/books\/(\d{1,})$`: http.HandlerFunc(b.find),
	}
}

func (b *bookApi) update(w http.ResponseWriter, r *http.Request) {
	resp, err := b.a.Update(r.Context(), nil)

	if err != nil {
		server.HttpError(w, err)
		return
	}

	server.Created(w, resp)
}

func (b *bookApi) find(w http.ResponseWriter, r *http.Request) {
	group := urlpattern["/{id}"].FindStringSubmatch(r.URL.Path)
	if len(group) == 0 {
		server.HttpError(w, fmt.Errorf("invalid url"))
		return
	}

	// id := group[0]
	resp, err := b.a.Find(r.Context(), nil)

	if err != nil {
		server.HttpError(w, err)
		return
	}

	server.Ok(w, resp)
}

func (b *bookApi) getAll(w http.ResponseWriter, r *http.Request) {
	resp, err := b.a.FindAll(r.Context(), nil)

	if err != nil {
		server.HttpError(w, err)
		return
	}

	server.Ok(w, resp)
}

func (b *bookApi) create(w http.ResponseWriter, r *http.Request) {
	var data dto.ReqNewBook

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		server.HttpError(w, err)
		return
	}

	resp, err := b.a.Create(r.Context(), nil)

	if err != nil {
		server.HttpError(w, err)
		return
	}

	server.Created(w, resp)
}
