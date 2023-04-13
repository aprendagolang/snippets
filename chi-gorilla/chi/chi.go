package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/echo", func(r chi.Router) {
		r.Get("/{name}", ChiHelloWorld)
	})
	r.Mount("/health", func() http.Handler {
		nr := chi.NewRouter()
		nr.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
		return nr
	}())

	http.ListenAndServe(":3000", r)
}

func ChiHelloWorld(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Olá %s!!!", name),
	})
}
