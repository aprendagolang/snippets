package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Nome string
	CPF string
	Idade int
}

func main() {
	file, err := os.Open("rh.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	people := make([]Person, len(records))
	for k, record := range records {
		idade, _ := strconv.Atoi(record[2])
		person := Person{
			Nome: record[0],
			CPF: record[1],
			Idade: idade,
		}

		people[k] = person
	}

	fmt.Println(people)
}