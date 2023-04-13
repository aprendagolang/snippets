package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	next := now.Add(13 * 24 * time.Hour)
	fmt.Println(next)
	fmt.Println(next.Format("02/01/06 03:04:05"))

	manual := time.Date(now.Year(), now.Month(), now.Day()+13, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	fmt.Println(manual)

	tstr := "27/04/2022 07:34:45 -03"
	parsed, err := time.Parse("02/01/2006 03:04:05 -07", tstr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parsed)
}
