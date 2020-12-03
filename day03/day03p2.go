package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("Couldn't read file")
	}

	return strings.Split(string(data), "\n")
}

func countTrees(rows []string, right, down int) int {
	currentColumn := 0
	treesEncountered := 0
	for i := 0; i < len(rows); i += down {
		if rows[i][currentColumn:currentColumn+1] == "#" {
			treesEncountered++
		}
		currentColumn = (currentColumn + right) % len(rows[0])
	}
	return treesEncountered
}

func main() {
	rows := readFile("input.txt")
	values := make([]int, 0)
	values = append(values, countTrees(rows, 1, 1))
	values = append(values, countTrees(rows, 3, 1))
	values = append(values, countTrees(rows, 5, 1))
	values = append(values, countTrees(rows, 7, 1))
	values = append(values, countTrees(rows, 1, 2))
	fmt.Printf("multiply(%v) = %v", values, values[0]*values[1]*values[2]*values[3]*values[4])
}
