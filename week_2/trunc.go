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
		resultFloat, _ := strconv.ParseFloat(input, 64)
		resultInt := int(resultFloat)
		fmt.Printf("Entered number is %d", resultInt)
	}

}
