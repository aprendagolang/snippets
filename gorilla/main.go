package main

import (
	"net/http"

	"github.com/aprendagolang/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", handlers.HandleHello).Methods("GET")
	r.HandleFunc("/person", handlers.HandlePerson).Methods("POST")

	http.ListenAndServe(":8080", r)
}
