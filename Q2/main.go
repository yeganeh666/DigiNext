// Minimum number of deletions required to make "AB" frequencies
package main

import (
	"fmt"
	"log"
	"strings"
)

func getInput() string {
	var str string
	fmt.Println("Enter string...")
	if _, err := fmt.Scanln(&str); err != nil {
		log.Fatal(err)
	}
	str = strings.TrimSpace(str)
	return str
}

func main() {
	str := getInput()
	var numberOfDeletion int
	for i := 1; i < len(str); i++ {
		if str[i-1] == str[i] {
			numberOfDeletion++
		}
	}
	fmt.Println("Minimum number of deletion: ", numberOfDeletion)
}
