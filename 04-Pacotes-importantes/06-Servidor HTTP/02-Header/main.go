package main

import "net/http"

func main() {
	http.HandleFunc("/", HelloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!\n"))
}
