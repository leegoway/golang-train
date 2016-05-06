/*
github.com/codegangsta/Negroni体验
*/
package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	mux.HandleFunc("/image", func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Path[1:]
		fmt.Fprintf(w, "This is an %s", name)
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
