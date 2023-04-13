package main

import "fmt"

type Pessoa struct {
	Nome      string
	Documento string
}

type PessoaFisica struct {
	Pessoa
	Idade uint8
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
}

type Cache[V PessoaFisica | PessoaJuridica] struct {
	data map[string]V
}

func New[T PessoaFisica | PessoaJuridica]() Cache[T] {
	return Cache[T]{}
}

func (c *Cache[V]) Set(key string, value V) {
	if c.data == nil {
		c.data = make(map[string]V)
	}

	c.data[key] = value
}

func (c *Cache[V]) Get(key string) (v V) {
	if c.data == nil {
		return
	}

	if v, ok := c.data[key]; ok {
		return v
	}

	return
}

func main() {
	pf := PessoaFisica{
		Pessoa{"Tiago", "000.000.000-00"},
		32,
	}

	cachepf := Cache[PessoaFisica]{}
	cachepf.Set(pf.Documento, pf)

	pj := PessoaJuridica{
		Pessoa{"Aprenda Golang LTDA", "00.000.000/0000-00"},
		"Aprenda Golang",
	}

	cachepj := Cache[PessoaJuridica]{}
	cachepj.Set(pj.Documento, pj)

	fmt.Println(cachepf.Get("000.000.000-00"))
	fmt.Println(cachepf.Get("000.000.110-00"))
	fmt.Println("-------")
	fmt.Println(cachepj.Get("00.000.000/0000-00"))
	fmt.Println(cachepj.Get("00.000.000/0000-01"))
}
