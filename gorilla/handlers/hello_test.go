package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandleHello(t *testing.T) {
	// struct auxiliar para os testes
	tt := []struct {
		rVar     string
		expected string
	}{
		{rVar: "Tiago", expected: "Hello Tiago"},
		{rVar: "Lucas", expected: "Hello Lucas"},
		{rVar: "Maria", expected: "Hello Maria"},
		{rVar: "Dani", expected: "Hello Dani"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/%s", tc.rVar)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/{name}", HandleHello).Methods("GET")
		router.ServeHTTP(rr, req)

		// In this case, our MetricsHandler returns a non-200 response
		// for a route variable it doesn't know about.
		var response map[string]string
		json.Unmarshal(rr.Body.Bytes(), &response)

		if response["message"] != tc.expected {
			t.Errorf("wrong response body for param %s: got %v want %v",
				tc.rVar, rr.Body.String(), tc.expected)
		}
	}
}
