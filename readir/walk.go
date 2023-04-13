package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk("/tmp/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: nome: %s\n", info.IsDir(), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
