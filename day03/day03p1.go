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

func main() {
	rows := readFile("input.txt")
	currentColumn := 0
	treesEncountered := 0
	for i := 0; i < len(rows); i++ {
		if rows[i][currentColumn:currentColumn+1] == "#" {
			treesEncountered++
		}
		currentColumn = (currentColumn + 3) % len(rows[0])
	}
	fmt.Println(treesEncountered)
}
