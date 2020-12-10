package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

func getMagicNumber(numbers []int) int {
	sort.Ints(numbers)
	ones := 0
	threes := 0
	for i := 0; i < len(numbers); i++ {
		prior := 0
		if i > 0 {
			prior = numbers[i-1]
		}
		diff := numbers[i] - prior
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		} else {
			fmt.Println("other")
		}
	}
	fmt.Println(ones, " ", threes)
	return ones * (threes + 1)
}

func main() {
	numbers := readNumbers("input.txt")

	fmt.Println(getMagicNumber(numbers))
}
