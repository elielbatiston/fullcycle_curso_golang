package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	// Como usar
	// encoder := json.NewEncoder(os.Stdout).Encode(conta)
	// ou
	// encoder := json.NewEncoder(os.Stdout)
	// encoder.Encode(conta)
	// ou
	// json.NewEncoder(os.Stdout).Encode(conta)
	// ou
	err := json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}
}
