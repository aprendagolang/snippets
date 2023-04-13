package main

import (
	"fmt"
	"time"

	"github.com/aprendagolang/pv/person"
)

func main() {
	p := person.Person{
		Name:      "Tiago Temporin",
		BirthDate: time.Date(2005, 02, 22, 7, 0, 0, 0, time.Local),
	}

	person.YearPointer(&p)

	base := time.Now().AddDate(-18, 0, 0)
	if base.After(p.BirthDate) {
		fmt.Println("Mais de 18 anos")
	} else {
		fmt.Println("Menos de 18 anos")
	}
}
