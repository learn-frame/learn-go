package main

import (
	"fmt"

	"learn-go/api/calculator"
)

func main() {
	result := calculator.Sum(1, 1)
	fmt.Println(result)
}
