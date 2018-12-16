package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Checks if the passed error is not null:
func checkErrors(err error) {
	if err != nil {
		fmt.Println("Whoops!")
		fmt.Println(err)
		os.Exit(1)
	}
}

// Makes an http request and parses its response:
func makeRequest(adress string) string {
	resp, err := http.Get(adress)
	checkErrors(err)
	bytes, err := ioutil.ReadAll(resp.Body)
	checkErrors(err)
	resp.Body.Close()

	return string(bytes)
}

func main() {
	for i := 0; i <= 50; i++ {
		fmt.Println(fmt.Sprintf("Request %d...", i) + "\n")

		// Makes the request:
		resp := makeRequest(fmt.Sprintf("http://localhost:3000/%d", i))

		fmt.Println("Resposta:")
		fmt.Println(resp)

		// Do not print the dashes if we are in the last iteration:
		if i != 50 {
			fmt.Println(strings.Repeat("-", 50))
		}
	}
}
