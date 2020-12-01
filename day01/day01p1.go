package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readNumbers(file string) []int {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("Couldn't read file.")
	}

	fileContents := string(data)
	lines := strings.Split(fileContents, "\n")

	output := make([]int, len(lines))
	for i, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("Couldn't parse (%v) into an int", number))
		}
		output[i] = number
	}
	return output
}

func main() {
	numbers := readNumbers("input.txt")

	seen := make(map[int]bool)
	for _, number := range numbers {
		complement := 2020 - number
		_, found := seen[complement]
		if found {
			fmt.Println(number * complement)
			return
		}

		seen[number] = true
	}
}
