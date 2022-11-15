package main

import "encoding/json"

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contax Conta
	err := json.Unmarshal(jsonPuro, &contax)
	if err != nil {
		println(err)
	}
	println(contax.Saldo)
}
