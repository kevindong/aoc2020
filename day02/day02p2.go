package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readLines(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("Couldn't read file")
	}

	fileContents := string(data)
	return strings.Split(fileContents, "\n")
}

func isPasswordValid(line string) bool {
	components := strings.Split(line, " ")

	policyCharacter := components[1][:1]
	password := components[2]
	matches := 0

	numbers := strings.Split(components[0], "-")
	for _, rawString := range numbers {
		index, err := strconv.Atoi(rawString)
		if err != nil {
			panic(fmt.Sprintf("Couldn't convert: %v", index))
		}
		currentCharacter := password[index-1 : index]
		if currentCharacter == policyCharacter {
			matches++
		}
	}

	return matches == 1
}

func main() {
	lines := readLines("input.txt")
	validPasswords := 0
	for _, line := range lines {
		if isPasswordValid(line) {
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
}
