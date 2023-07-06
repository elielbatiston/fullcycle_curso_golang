package main

import (
	"fmt"
)

func main() {
	var minhaVar interface{} = "Eliel Batiston"
	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	fmt.Printf("=========================\n")
	fmt.Printf("Agora vai dar um panic porque nao tratamos erro\n")

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
