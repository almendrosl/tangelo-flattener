package v1

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"tangelo-flattener/internal/data"
	"tangelo-flattener/pkg/flatteners"
	"tangelo-flattener/pkg/response"
)

type OutputResponse struct {
	Flatten []interface{} `json:"flatten"`
	Depth   int           `json:"depth"`
}

func New() http.Handler {
	r := chi.NewRouter()

	r.Post("/flatArray", flatArrayHandler)

	r.Get("/flatArray", getListSuccessfulProcessedArray)

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

	var arrayToFlatten flatteners.FirstAlgorithm = inputArray.([]interface{})

	outputArray, depth := arrayToFlatten.FlattenArray()

	if err = data.SaveData(inputArray.([]interface{}), outputArray); err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	_ = response.JSON(w, r, http.StatusOK, OutputResponse{Flatten: outputArray, Depth: depth})
}

func getListSuccessfulProcessedArray(w http.ResponseWriter, r *http.Request) {
	array, err := data.GetListSuccessfulProcessedArray(100)
	if err != nil {
		_ = response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	_ = response.JSON(w, r, http.StatusOK, array)
}
