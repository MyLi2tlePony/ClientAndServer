package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n        string
		elements string
	)

	in := bufio.NewReader(os.Stdin)
	n, _ = in.ReadString('\n')
	n = strings.TrimSuffix(n, "\r\n")

	elements, _ = in.ReadString('\n')
	elements = strings.TrimSuffix(elements, "\n")

	elementsNumber := ConvertToInt(n)
	numbers := ConvertToArrayInt(&elements)

	if !IsSort(numbers) {
		fmt.Println(-1)
	} else {
		fmt.Println(numbers[elementsNumber-1] - numbers[0])
	}
}

func IsSort(numbers []int) bool {
	for index := 1; index < len(numbers); index++ {
		if numbers[index] < numbers[index-1] {
			return false
		}
	}
	return true
}

func ConvertToInt(number string) int {

	elementsNumber, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return elementsNumber
}

func ConvertToArrayInt(text *string) []int {
	elements := strings.Fields(*text)
	var numbers []int

	for _, element := range elements {
		number, err := strconv.Atoi(element)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		numbers = append(numbers, number)
	}
	return numbers
}
