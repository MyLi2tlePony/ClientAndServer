package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Id       int    `json:"id"`
	UserName string `json:"user-name"`
	Text     User   `json:"text"`
}

func main() {
	users := []User{
		{1, "Andrey"},
		{2, "Dima"},
	}

	http.HandleFunc("/get/user/all", func(response http.ResponseWriter, request *http.Request) {
		enc := json.NewEncoder(response)
		if err := enc.Encode(users); err != nil {
			fmt.Printf("error encoding struct into JSON: %v\n", err)
		}
	})

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Index Page")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
