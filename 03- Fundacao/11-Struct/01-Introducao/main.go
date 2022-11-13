package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	cli := Cliente{
		Nome:  "Eliel",
		Idade: 30,
		Ativo: true,
	}
	cli.Ativo = false

	fmt.Println(cli)
}
