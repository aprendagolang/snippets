package main

import (
	_ "embed"
	"fmt"
)

//go:embed static.txt
var content string

func main() {
	fmt.Println(content)
}
