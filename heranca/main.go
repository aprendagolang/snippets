package main

import (
	"fmt"
	"time"
)

type Pessoa struct {
	Nome           string
	Documento      string
	DataNascimento time.Time
	Enderecos      []Endereco
}

func (p *Pessoa) Idade() int {
	return time.Now().Year() - p.DataNascimento.Year()
}

type PessoaJuridica struct {
	Pessoa
	InscricaoEstadual string
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	RG        string
}

func (pf *PessoaFisica) NomeCompleto() string {
	return fmt.Sprintf("%s %s", pf.Nome, pf.Sobrenome)
}

type Endereco struct {
	Tipo     string
	Endereco string
	Bairro   string
	Cidade   string
	Estado   string
}

func (e *Endereco) Completo() string {
	return fmt.Sprintf("%s: %s, %s, %s - %s", e.Tipo, e.Endereco, e.Bairro, e.Cidade, e.Estado)
}

func main() {
	pf := PessoaFisica{
		Pessoa{
			Nome:           "Tiago",
			Documento:      "000.000.000-00",
			DataNascimento: time.Date(1990, 02, 22, 7, 0, 0, 0, time.Local),
			Enderecos: []Endereco{
				{
					Tipo:     "casa",
					Endereco: "Rua dos Gophers",
					Bairro:   "Imperial",
					Cidade:   "GoCity",
					Estado:   "GOstate",
				},
			},
		},
		"Temporin",
		"00.000.000-00",
	}

	fmt.Println("Nome:", pf.NomeCompleto())
	fmt.Println("Idade:", pf.Idade())
	for _, e := range pf.Enderecos {
		fmt.Println("Endereço:", e.Completo())
	}
}
