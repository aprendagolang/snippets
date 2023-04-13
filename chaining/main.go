package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     uint
}

func (p *Pessoa) SetNome(nome string) *Pessoa {
	p.nome = nome
	return p
}

func (p *Pessoa) SetSobrenome(sobrenome string) *Pessoa {
	p.sobrenome = sobrenome
	return p
}

func (p *Pessoa) SetIdade(idade uint) *Pessoa {
	p.idade = idade
	return p
}

func (p *Pessoa) Print() {
	fmt.Printf("Olá, meu nome é %s %s e eu tenho %d anos.", p.nome, p.sobrenome, p.idade)
}

func main() {
	p := Pessoa{}
	p.SetNome("Tiago").
		SetSobrenome("Temporin").
		SetIdade(32).
		Print()
}
