package main

import "fmt"

func main() {
	salarios := map[string]int{"Eliel": 1000, "João": 2000, "Maria": 3000}
	delete(salarios, "Eliel")
	salarios["Fulano"] = 5000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

	fmt.Printf("============\n")

	sal := make(map[string]int)
	sal["Eliel"] = 1000

	fmt.Printf("O salario é %d\n", sal["Eliel"])

	fmt.Printf("============\n")

	sal1 := map[string]int{}
	sal1["Eliel"] = 2000

	fmt.Printf("O salario é %d\n", sal1["Eliel"])
}
