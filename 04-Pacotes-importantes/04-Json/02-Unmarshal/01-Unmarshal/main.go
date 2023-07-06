package main

import "encoding/json"

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var contax Conta
	err := json.Unmarshal(jsonPuro, &contax)
	if err != nil {
		println(err)
	}
	println(contax.Saldo)
}
