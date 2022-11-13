package main

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c Conta) simularSemPonteiro(valor int) int {
	c.saldo += valor
	println("Novo saldo = ", c.saldo)
	return c.saldo
}

func (c *Conta) simularComPonteiro(valor int) int {
	c.saldo += valor
	println("Novo saldo = ", c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}

	conta.simularSemPonteiro(200)
	println("Saldo = ", conta.saldo)

	println("===============")

	conta.simularComPonteiro(200)
	println("Saldo = ", conta.saldo)
}
