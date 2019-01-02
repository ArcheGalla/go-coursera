package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Enter a float number please :")

	_, isError := fmt.Scanln(&input)

	if isError == nil {
		result, _ := strconv.ParseFloat(input, 64)
		fmt.Println("Entered number is", result)
	}

}
