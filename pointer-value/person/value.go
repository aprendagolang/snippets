package person

func YearValue(p Person) {
	p.BirthDate = p.BirthDate.AddDate(-1, 0, 0)
}
