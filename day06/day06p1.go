package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readGroups(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("Can't read file.")
	}

	return strings.Split(string(data), "\n\n")
}

func count(answers string) int {
	set := make(map[string]bool)
	for _, c := range answers {
		set[string(c)] = true
	}
	delete(set, "\n")
	return len(set)
}

func main() {
	groups := readGroups("input.txt")
	counter := 0
	for _, group := range groups {
		counter += count(group)
	}
	fmt.Println(counter)
}
