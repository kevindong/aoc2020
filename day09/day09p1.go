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
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	numbers := make([]int, len(lines))
	for i, number := range lines {
		integer, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		numbers[i] = integer
	}
	return numbers
}

func getAnamoly(preamble int, numbers []int) int {
	window := make(map[int]int)
	for i := 0; i < preamble; i++ {
		window[numbers[i]]++
	}

	for i := preamble; i < len(numbers); i++ {
		if !isSumPresent(numbers[i], window) {
			return numbers[i]
		}

		window[numbers[i-preamble]]--
		if window[numbers[i-preamble]] == 0 {
			delete(window, numbers[i-preamble])
		}
		window[numbers[i]]++
	}

	// There was a problem
	return -1
}

func isSumPresent(sum int, window map[int]int) bool {
	for key := range window {
		complement := sum - key
		if complement != key && window[complement] > 0 {
			return true
		} else if complement == key && window[complement] >= 2 {
			return true
		}
	}
	return false
}

func main() {
	numbers := readNumbers("input.txt")
	fmt.Println(getAnamoly(25, numbers))
}
