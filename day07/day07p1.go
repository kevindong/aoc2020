package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getCount(file string) int {
	parents := make(map[string][]string)

	lines := readLines(file)
	for _, line := range lines {
		// Get the current bag description
		parentDescription := strings.Split(line, " contain ")[0]

		// Get the bags layered within and insert
		contents := strings.Split(line, " contain ")[1]
		contentsSplit := strings.Split(contents, ", ")
		for _, child := range contentsSplit {
			if child == "no other" {
				continue
			}
			childDescription := strings.TrimLeft(child, "0123456789 ")
			parents[childDescription] = append(parents[childDescription], parentDescription)
		}
	}

	possibilities := make(map[string]bool)
	fill(parents, possibilities, "shiny gold")
	return len(possibilities)
}

func fill(parents map[string][]string, output map[string]bool, bag string) {
	_, found := output[bag]
	if len(parents[bag]) == 0 || found {
		return
	}
	for _, possibility := range parents[bag] {
		fill(parents, output, possibility)
		output[possibility] = true
	}
}

func readLines(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	filtered := strings.ReplaceAll(string(data), " bags", "")
	filtered = strings.ReplaceAll(filtered, " bag", "")
	filtered = strings.ReplaceAll(filtered, ".", "")
	return strings.Split(filtered, "\n")
}

func main() {
	fmt.Println(getCount("input.txt"))
}
