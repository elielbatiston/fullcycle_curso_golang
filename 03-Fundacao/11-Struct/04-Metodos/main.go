package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {
	cli := Cliente{
		Nome:  "Eliel",
		Idade: 30,
		Ativo: true,
	}
	cli.Idade = 35

	fmt.Println(cli)
	fmt.Println("=============")

	cli.Desativar()

	fmt.Println("=============")
	fmt.Println("Mas continua com os mesmos valores")
	fmt.Println(cli)
}
