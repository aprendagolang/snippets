package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/{name}", func(rw http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(map[string]string{
			"message": fmt.Sprintf("Olá %s!!!", name),
		})
	})

	http.ListenAndServe(":8080", r)
}
