package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	// messages
	found := "Found!"
	notFount := "Not Found!"

	// search criteria
	start := "i"
	contains := "a"
	end := "n"

	fmt.Println("Please, enter the string ")
	_, _ = fmt.Scanln(&input)
	lowerInput := strings.ToLower(input)

	if strings.HasPrefix(lowerInput, start) && strings.HasSuffix(lowerInput, end) && strings.Contains(lowerInput, contains) {
		fmt.Println(found)
	} else {
		fmt.Println(notFount)
	}

}
