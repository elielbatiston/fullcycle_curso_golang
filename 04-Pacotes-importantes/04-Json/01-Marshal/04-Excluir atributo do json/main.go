package main

import "encoding/json"

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"-"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(res)
	println(string(res))
}
