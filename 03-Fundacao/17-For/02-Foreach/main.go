package main

func main() {
	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}
}
