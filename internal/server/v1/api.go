package v1

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"tangelo-flattener/pkg/response"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Post("/flatArray", flatArrayHandler)

	return r
}

func flatArrayHandler(w http.ResponseWriter, r *http.Request) {
	var inputArray interface{}
	err := json.NewDecoder(r.Body).Decode(&inputArray)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	_ = response.JSON(w, r, http.StatusCreated, inputArray)
}
