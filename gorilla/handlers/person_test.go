package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandlePerson(t *testing.T) {
	tt := []struct {
		rVar     Person
		expected string
	}{
		{rVar: Person{Name: "Tiago", Age: 31}, expected: "Nome: Tiago - Idade: 31\n"},
		{rVar: Person{Name: "Lucas", Age: 41}, expected: "Nome: Lucas - Idade: 41\n"},
		{rVar: Person{Name: "Maria", Age: 21}, expected: "Nome: Maria - Idade: 21\n"},
		{rVar: Person{Name: "Dani", Age: 36}, expected: "Nome: Dani - Idade: 36\n"},
	}

	for _, tc := range tt {
		data, err := json.Marshal(tc.rVar)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/person", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/person", HandlePerson).Methods("POST")
		router.ServeHTTP(rr, req)

		// In this case, our MetricsHandler returns a non-200 response
		// for a route variable it doesn't know about

		if rr.Body.String() != tc.expected {
			t.Errorf("wrong response body for param %v: got %v want %v",
				tc.rVar, rr.Body.String(), tc.expected)
		}
	}
}
