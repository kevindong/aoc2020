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

	password := components[2]

	numbers := strings.Split(components[0], "-")
	firstIndex, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(fmt.Sprintf("Couldn't convert: %v", numbers[0]))
	}
	firstIndexCharacter := string(password[firstIndex-1])
	secondIndex, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(fmt.Sprintf("Couldn't convert: %v", numbers[1]))
	}
	secondIndexCharacter := string(password[secondIndex-1])

	character := string(components[1][0])

	return (firstIndexCharacter == character && secondIndexCharacter != character) || (firstIndexCharacter != character && secondIndexCharacter == character)
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
