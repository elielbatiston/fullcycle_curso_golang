package main

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
func main() {
	forever := make(chan bool) //Canal vazio

	forever <- true //Canal cheio, porem, como o canal foi alterado dentro da propria thread 1 (thread principal), vai dar deadlock. Para não dar deadlock, o canal tem que ficar cheio em outra go routines

	<-forever //esta esperando o canal ficar cheio para ele esvaziar. Com isso, ele vai segurar o processo de pe do nosso main
}
