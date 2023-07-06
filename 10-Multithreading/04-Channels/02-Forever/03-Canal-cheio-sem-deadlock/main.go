package main

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
func main() {
	forever := make(chan bool) //Canal vazio

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true //Canal cheio
	}()

	<-forever //esta esperando o canal ficar cheio para ele esvaziar. Com isso, ele vai segurar o processo de pe do nosso main
}
