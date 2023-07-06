package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

func main() {
	cli := Cliente{
		Nome:  "Eliel",
		Idade: 30,
		Ativo: true,
		Address: Endereco{
			Logradouro: "Rua X",
			Numero:     1,
			Cidade:     "Cidade X",
			Estado:     "Estado X",
		},
	}
	cli.Ativo = false

	fmt.Println(cli)
	fmt.Println("==============")
	fmt.Println("Utilizacao de atributos de struc")
	fmt.Println(cli.Address.Logradouro)
	fmt.Println("Assim nao funcion [cli.Cidade]")
	//fmt.Println(cli.Cidade)
}
