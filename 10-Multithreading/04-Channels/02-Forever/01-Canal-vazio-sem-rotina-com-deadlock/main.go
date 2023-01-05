package main

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
func main() {
	forever := make(chan bool) //Canal vazio

	<-forever //esta esperando o canal ficar cheio para ele esvaziar. Com isso, ele vai segurar o processo de pe do nosso main

	//Se eu manter só assim e executar o programa vou ter deadlock
}
