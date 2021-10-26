package logs

import (
	"log"
	"os"
)

// Write escreve uma mensagem de log em um arquivo
func Write(message, filepath string) {
	// define um caminho se o filepath não for passado
	if filepath == "" {
		filepath = "/tmp/my-app.log"
	}

	// abre o arquivo onde vamos escrever nossa mensagem
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	// seta a saída do log para o arquivo que abrimos
	log.SetOutput(file)

	// loga data e hora
	log.SetFlags(log.LstdFlags)

	// escreve a mensagem
	log.Println(message)
}
