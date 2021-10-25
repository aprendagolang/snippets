package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// struct para mapear as configs do arquivo config.toml
type Config struct {
	Host string `toml:"host"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
	Db   string `toml:"db"`
}

// variável onde as configs serão armazenadas
var conf Config

// função que é executada antes do programa iniciar
func init() {
	// leitura do arquivos config.toml
	data, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}

	// decode do conteúdo do arquivo para a struct Config
	if _, err := toml.Decode(string(data), &conf); err != nil {
		panic(err)
	}
}

func main() {
	// print para validação
	fmt.Printf("Host: %s\n", conf.Host)
	fmt.Printf("User: %s\n", conf.User)
	fmt.Printf("Pass: %s\n", conf.Pass)
	fmt.Printf("Db: %s\n", conf.Db)
}
