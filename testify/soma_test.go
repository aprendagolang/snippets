package soma_test

import (
	"testing"

	"github.com/aprendagolang/soma"
	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	comNill := make([]int, 3)
	comNill[0] = 10
	comNill[1] = 10

	testes := []struct {
		Valores   []int
		Resultado int
	}{
		{Valores: []int{1, 2, 3}, Resultado: 6},
		{Valores: []int{1, 2, 3, 4}, Resultado: 10},
		{Valores: []int{3, 3, 3, 3}, Resultado: 12},
		{Valores: []int{1, 1, 1, 1}, Resultado: 4},
		{Valores: []int{12, 20, 35}, Resultado: 67},
		{Valores: []int{19, 21, 32}, Resultado: 72},
		{Valores: comNill, Resultado: 20},
	}

	for _, teste := range testes {
		total := soma.Soma(teste.Valores...)
		assert.Equal(t, total, teste.Resultado, "Valor esperado: %d - Valor retornado: %d", teste.Resultado, total)
	}
}
