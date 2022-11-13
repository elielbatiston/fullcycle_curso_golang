package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	cli := Cliente{
		Nome:  "Eliel",
		Idade: 30,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua X",
			Numero:     1,
			Cidade:     "Cidade X",
			Estado:     "Estado X",
		},
	}
	cli.Ativo = false

	fmt.Println(cli)
	fmt.Println("==============")
	fmt.Println("Utilizacao da composicao")
	fmt.Println(cli.Logradouro)
	fmt.Println("Ou")
	fmt.Println(cli.Endereco.Cidade)
}
