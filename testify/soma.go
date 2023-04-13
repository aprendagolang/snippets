package soma

// Soma soma N valores e retorna seu total
func Soma(valores ...int) (total int) {
	for _, valor := range valores {
		total += valor
	}

	return
}
