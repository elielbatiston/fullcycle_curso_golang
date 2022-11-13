package main

func soma(a int, b *int) int {
	a = a * 2
	*b = *b * 3
	return a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	println("minhaVar1 antes de soma = ", minhaVar1)
	println("minhaVar2 antes de soma = ", minhaVar2)
	println("=================")

	println("Soma = ", soma(minhaVar1, &minhaVar2))

	println("minhaVar1 depois de soma = ", minhaVar1)
	println("minhaVar2 depois de soma = ", minhaVar2)
}
