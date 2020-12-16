package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(file string) []int {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	stringSlice := strings.Split(string(data), ",")
	intSlice := make([]int, 0)
	for _, numberString := range stringSlice {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}

		intSlice = append(intSlice, number)
	}
	return intSlice
}

func getMagicNumber(numbers []int) int {
	lastSeen := make(map[int]int)
	for i, number := range numbers {
		fmt.Println(number)
		lastSeen[number] = i + 1
	}

	lastNumber := numbers[len(numbers)-1]
	for i := len(numbers) + 1; i <= 2020; i++ {
		save := lastNumber
		prior := lastSeen[lastNumber]
		if prior == 0 {
			lastNumber = 0
		} else {
			lastNumber = i - prior - 1
		}
		lastSeen[save] = i - 1
	}
	return lastNumber
}

func main() {
	input := readInput("input.txt")
	fmt.Println(getMagicNumber(input))
}
