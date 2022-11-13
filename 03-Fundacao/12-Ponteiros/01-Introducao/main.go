package main

func main() {
	// Memória -> Endreço -> Valor
	a := 10
	println("Ponteiro de a = ", &a)
	println("Valor de a = ", a)

	var ponteiro *int = &a
	println("====================")
	println("Endereco de ponteiro = ", ponteiro)
	println("Valor de ponteiro = ", *ponteiro)

	*ponteiro = 20
	println("====================")
	println("Endereco de ponteiro = ", ponteiro)
	println("Valor de ponteiro = ", *ponteiro)
	println("Valor de a = ", a)

	b := &a
	*b = 30
	println("====================")
	println("Endereco de b = ", b)
	println("Valor de b = ", *ponteiro)
	println("Valor de a = ", a)
}
