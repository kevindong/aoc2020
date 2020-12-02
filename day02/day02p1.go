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

	numbers := strings.Split(components[0], "-")
	minValue, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(fmt.Sprintf("Couldn't convert: %v", numbers[0]))
	}
	maxValue, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(fmt.Sprintf("Couldn't convert: %v", numbers[1]))
	}

	character := string(components[1][0])

	password := components[2]

	countInPassword := strings.Count(password, character)
	return minValue <= countInPassword && countInPassword <= maxValue
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
