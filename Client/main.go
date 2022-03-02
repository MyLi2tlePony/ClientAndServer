package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Чат открыт, можно писать...")

	for {
		userMessage := ConsoleReadLn()

		switch userMessage {
		case "users":
			GetAllUsers()
		}
	}
}

func GetAllUsers() {
	httpURL := "http://localhost:8181/get/user/all"
	httpMethod := "POST"
	client := &http.Client{}

	req, _ := http.NewRequest(
		httpMethod, httpURL, nil,
	)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func ConsoleReadLn() string {
	consoleReader := bufio.NewReader(os.Stdin)
	message, _ := consoleReader.ReadString('\n')

	message = strings.TrimSuffix(message, "\r")
	message = strings.TrimSuffix(message, "\n")

	return message
}
