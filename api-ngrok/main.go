package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		io.WriteString(rw, "Olá Mundo!")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
