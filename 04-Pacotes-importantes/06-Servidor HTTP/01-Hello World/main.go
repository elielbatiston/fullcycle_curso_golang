package main

import "net/http"

func main() {
	http.HandleFunc("/", HelloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}
