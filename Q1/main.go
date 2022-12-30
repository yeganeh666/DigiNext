// Count of 'a' in a repeated string
package main

import (
	"fmt"
	"log"
	"strings"
)

var subString = "a"

func getInput() (string, int) {
	var str string
	fmt.Println("Enter string...")
	if _, err := fmt.Scanln(&str); err != nil {
		log.Fatal(err)
	}
	str = strings.TrimSpace(str)

	var length int
	fmt.Println("Enter length...")
	if _, err := fmt.Scanln(&length); err != nil {
		log.Fatal(err)
	}

	return str, length
}
func main() {
	str, length := getInput()
	repeatedString := strings.Repeat(str, length)
	fmt.Println("count of 'a' :", strings.Count(repeatedString[:length], subString))
}
