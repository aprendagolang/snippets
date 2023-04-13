package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func HandlePerson(w http.ResponseWriter, rq *http.Request) {
	var p Person

	err := json.NewDecoder(rq.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Nome: %s - Idade: %d\n", p.Name, p.Age)
}
