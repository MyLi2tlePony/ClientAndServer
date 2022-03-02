package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

}

func ConsoleReadLn() string {
	consoleReader := bufio.NewReader(os.Stdin)
	message, _ := consoleReader.ReadString('\n')

	message = strings.TrimSuffix(message, "\r")
	message = strings.TrimSuffix(message, "\n")

	return message
}
