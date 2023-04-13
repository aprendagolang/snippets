package main

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed template.yaml
var content string

type Person struct {
	Name string
	Age  string
}

func main() {
	if len(os.Args) < 2 {
		panic("min args 2")
	}

	p := Person{
		Name: os.Args[1],
		Age:  os.Args[2],
	}

	tmpl, err := template.New("person").Parse(content)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("output.yaml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, p)
	if err != nil {
		panic(err)
	}
}
