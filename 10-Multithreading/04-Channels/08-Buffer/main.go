package main

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
func main() {
	ch := make(chan string, 2) //Buffer de 2
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
