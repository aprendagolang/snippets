package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

var (
	people = [][]string{
		[]string{"1", "Tiago Temporin"},
		[]string{"2", "João Silva"},
		[]string{"3", "Mateus Cardoso"},
		[]string{"4", "Maria Lina"},
		[]string{"5", "Camila Manga"},
		[]string{"6", "Joice Santos"},
		[]string{"7", "Lucas Leal"},
		[]string{"8", "Vanessa da Terra"},
		[]string{"9", "Mateus de Morais"},
		[]string{"10", "Maria Luiza"},
	}

	orders = [][]string{
		[]string{"1", "5"},
		[]string{"2", "10"},
		[]string{"3", "0"},
		[]string{"4", "0"},
		[]string{"5", "2"},
		[]string{"6", "9"},
		[]string{"7", "3"},
		[]string{"8", "15"},
		[]string{"9", "3"},
		[]string{"10", "7"},
	}
)

func getFullName(id string) (name string) {
	for _, row := range people {
		if id == row[0] {
			name = row[1]
		}
	}

	return
}

func getTotalOrders(id string) (qtd string) {
	for _, row := range orders {
		if id == row[0] {
			qtd = row[1]
		}
	}

	return
}

func main() {
	r := mux.NewRouter()
	r.Use(middleware)

	r.HandleFunc("/cc/{id}", NameOrders)
	r.HandleFunc("/sc/{id}", NameOrders)

	http.ListenAndServe(":8000", r)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var (
			vars   = mux.Vars(r)
			name   string
			orders string
		)

		if strings.Contains(r.URL.Path, "cc") {
			cname := make(chan string, 1)
			corders := make(chan string, 1)

			wg := sync.WaitGroup{}
			wg.Add(2)

			go func() {
				cname <- getFullName(vars["id"])
				wg.Done()
			}()

			go func() {
				corders <- getTotalOrders(vars["id"])
				wg.Done()
			}()

			wg.Wait()

			name = <-cname
			orders = <-corders
		} else {
			name = getFullName(vars["id"])
			orders = getTotalOrders(vars["id"])
		}

		ctx := context.WithValue(r.Context(), "name", name)
		ctx = context.WithValue(ctx, "orders", orders)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

func NameOrders(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value("name")
	orders := r.Context().Value("orders")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s - Total orders: %s\n", name, orders)
}
