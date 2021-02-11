package v1

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetListSuccessfulProcessedArray(t *testing.T) {
	req, err := http.NewRequest("Get", "/flatArray", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getListSuccessfulProcessedArray)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test(t *testing.T) {
	var jsonStr = []byte(`[[2634, 9867, true, false, ["hola"]], ["chau"]]`)

	req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(flatArrayHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"flatten":[2634,9867,true,false,"hola","chau"],"depth":3}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
