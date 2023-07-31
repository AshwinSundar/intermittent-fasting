package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Enter time of first meal today (HHMM): ")
	var time string = readTime()
	validateTime(time)
}

func readTime() string {
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("Error reading input: ", err)
		return ""
	}	

	return input
}

func validateTime(time string) {
	pattern := "^[0-2][0-9][0-5][0-9]$"

	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex: ", err)
	}	

	match := regex.FindString(time)

	if match != "" {
		fmt.Printf("Format is correct: %s\n", match)
	} else {
		fmt.Println("Invalid time. Format must be HHMM")
	}
}
