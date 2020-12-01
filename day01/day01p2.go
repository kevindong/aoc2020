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

	set := make(map[int]bool, len(numbers))
	for _, number := range numbers {
		set[number] = true
	}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			complement := 2020 - numbers[i] - numbers[j]
			_, found := set[complement]
			if found {
				output := numbers[i] * numbers[j] * complement
				fmt.Printf("%v * %v * %v = %v", numbers[i], numbers[j], complement, output)
				return
			}
		}
	}
}
