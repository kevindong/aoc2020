package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getCount(file string) int {
	children := make(map[string][]string)

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
			childNumber, err := strconv.Atoi(strings.Split(child, " ")[0])
			if err != nil {
				panic(err)
			}

			for i := 0; i < childNumber; i++ {
				children[parentDescription] = append(children[parentDescription], childDescription)
			}
		}
	}

	cache := make(map[string]int)

	return countBags(children, cache, "shiny gold")
}

func countBags(children map[string][]string, cache map[string]int, bag string) int {
	count, found := cache[bag]
	if found {
		return count
	} else if len(children[bag]) == 0 {
		return 0
	}

	counter := 0
	for _, child := range children[bag] {
		result := countBags(children, cache, child)
		cache[child] = result
		counter += 1 + result
	}
	return counter
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
