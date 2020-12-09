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

	panic("Anamoly number not found")
}

func findMagicNumber(sum int, numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			windowSum := 0
			min := numbers[i]
			max := numbers[i]
			for k := i; k < j; k++ {
				windowSum += numbers[k]

				if numbers[k] < min {
					min = numbers[k]
				}
				if numbers[k] > max {
					max = numbers[k]
				}
			}
			// Not it
			if windowSum != sum {
				continue
			}

			// Found it; do some processing
			return min + max
		}
	}

	panic("Magic number not found")
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
	target := getAnamoly(25, numbers)
	fmt.Println(findMagicNumber(target, numbers))
}
