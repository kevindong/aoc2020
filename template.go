package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readInput(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func main() {
	lines := readInput("sample.txt")
	fmt.Println(lines)
}
