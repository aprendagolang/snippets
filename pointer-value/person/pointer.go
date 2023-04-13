package person

func YearPointer(p *Person) {
	p.BirthDate = p.BirthDate.AddDate(-1, 0, 0)
}
