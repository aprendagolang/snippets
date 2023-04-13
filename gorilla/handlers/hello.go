package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleHello(w http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Hello %s", vars["name"]),
	})
}
