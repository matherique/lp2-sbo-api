package http

import (
	"encoding/json"
	"net/http"

	"github.com/matherique/lp2-sbo-api/pkg/errors"
)

func Created(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func Ok(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func HttpError(w http.ResponseWriter, err error) {
	enc := json.NewEncoder(w)

	switch e := err.(type) {
	case errors.HttpError:
		w.WriteHeader(e.Status)
		enc.Encode(e)

	default:
		enc.Encode(errors.NewHttp(
			http.StatusInternalServerError,
			"internal server error",
		))
	}
}
