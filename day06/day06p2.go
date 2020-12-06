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
	numberOfPeople := len(strings.Split(answers, "\n"))

	set := make(map[string]int)
	for _, c := range answers {
		set[string(c)] = set[string(c)] + 1
	}
	delete(set, "\n")

	counter := 0
	for _, value := range set {
		if value == numberOfPeople {
			counter++
		}
	}

	return counter
}

func main() {
	groups := readGroups("input.txt")
	counter := 0
	for _, group := range groups {
		counter += count(group)
	}
	fmt.Println(counter)
}
