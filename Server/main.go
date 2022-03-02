package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/about", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "About Page")
	})
	http.HandleFunc("/contact", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Contact Page")
	})
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Index Page")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
