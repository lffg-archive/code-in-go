package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// This will be printed on the terminal.
	fmt.Printf("%s on %s\n", r.Method, r.URL.Path)

	// This will be printed on the webpage:
	fmt.Fprint(w, "Hello, world! From a Go server. :)")
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":3000", nil)
}
