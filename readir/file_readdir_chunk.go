package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("/tmp")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		files, err := f.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			continue
		}

		fmt.Println(files[0].Name(), files[0].IsDir())
	}
}
