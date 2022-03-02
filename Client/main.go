package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Чат открыт, можно писать...")

	userMessage := ConsoleReadLn()

	jsonData, _ := json.Marshal(map[string]string{
		"name":    "Toby",
		"message": userMessage,
	})

	httpURL := "http://localhost:8181/"
	httpMethod := "POST"
	client := &http.Client{}

	req, _ := http.NewRequest(
		httpMethod, httpURL, bytes.NewBuffer(jsonData),
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
