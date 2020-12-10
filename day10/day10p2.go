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
	dp := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		currentNumber := numbers[i]
		dp[currentNumber] = dp[currentNumber-1] + dp[currentNumber-2] + dp[currentNumber-3]
		if currentNumber <= 3 {
			dp[currentNumber]++
		}
	}
	return dp[numbers[len(numbers)-1]]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	numbers := readNumbers("input.txt")
	fmt.Println(getMagicNumber(numbers))
}
