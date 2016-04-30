package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/world", handleFunction)
	http.ListenAndServe(":8080", nil)
}

func handleFunction(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	fmt.Fprint(w, "Hello, "+name)
}
