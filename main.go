package main

import (
	"fmt"
	"os"
)

func main() {
	numbers := []Number{}
	for _, arg := range os.Args[1:] {
		numbers = append(numbers, New(arg))
	}

	for _, number := range numbers {
		fmt.Println(number, number.Score())
	}
}
