package main

import "fmt"

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
	fmt.Printf("Empresa %s\n", e.Nome)
}

type Funcionario struct {
	Nome string
}

func (f Funcionario) Desativar(value bool) {
	fmt.Printf("Funcionario %s, valor %v\n", f.Nome, value)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	cli := Cliente{
		Nome:  "Eliel",
		Idade: 30,
		Ativo: true,
	}

	cli.Desativar()
	Desativacao(cli)

	emp := Empresa{
		Nome: "XXXX",
	}

	emp.Desativar()
	Desativacao(emp)

	fun := Funcionario{
		Nome: "Jose",
	}

	fun.Desativar(true)
	//Desativacao(fun)
	println("Como funcionario nao tem a classe desativar com a mesma assinatura, nao podemos usar a funcao Desativacao(fun)")
}
