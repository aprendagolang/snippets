package main

var (
	pessoas = []Pessoa{
		{Nome: "Dani", Idade: 35},
		{Nome: "Lucas", Idade: 29},
		{Nome: "Rafael", Idade: 38},
	}
)

type Pessoa struct {
	Nome  string
	Idade int32
}

func Append1() {
	nomes := []string{"Tiago", "Alexandre", "Luara"}
	for _, pessoa := range pessoas {
		nomes = append(nomes, pessoa.Nome)
	}
}

func Append2() {
	nomes := []string{"Tiago", "Alexandre", "Luara"}

	todos := make([]string, 0, len(nomes)+len(pessoas))
	for _, nome := range nomes {
		todos = append(todos, nome)
	}

	for _, pessoa := range pessoas {
		todos = append(todos, pessoa.Nome)
	}
}

func Append3() {
	nomes := []string{"Tiago", "Alexandre", "Luara"}

	// o tamanho do array é a posição seguinte
	index := len(nomes)

	nomes = append(nomes, make([]string, len(pessoas))...)

	for _, pessoa := range pessoas {
		nomes[index] = pessoa.Nome
		index++
	}
}
