package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("subscribers.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	records := [][]string{
		{"Tiago Temporin", "tiago.temporin@provedor.com"},
		{"Amanda Moreira", "amanda.moreira@teste.io"},
		{"João Santos", "joao.santos@email.me"},
		{"Valentina Silva", "valentina.silva@tutorial.com"},
	}

    err = csv.NewWriter(file).WriteAll(records)
    if err != nil {
        panic(err)
    }
    /*
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}
    */
}