package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", GorillaHelloWorld).Methods("GET")

	http.ListenAndServe(":3000", r)
}

func GorillaHelloWorld(w http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Olá %s!!!", vars["name"]),
	})
}
