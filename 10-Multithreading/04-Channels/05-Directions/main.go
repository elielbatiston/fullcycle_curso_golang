package main

import "fmt"

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}

func recebe(nome string, hello chan<- string) {
	//Como essa função só recebe dado, eu posso setar a direção desse canal como "publish" com a marcação <- atras de chan
	//Isso significa que nosso channel apenas recebe informação
	//Canal chamado de receive only
	hello <- nome
}

func ler(data <-chan string) {
	//Como essa função só le o dado do canal (esvazia o canal) então eu posso setar a direçao desse canal para leitura com a marcação <- na frente do chan
	//Isso significa que nosso channel apenas le a informação
	//Canal chamado de send only
	fmt.Println(<-data)
}
