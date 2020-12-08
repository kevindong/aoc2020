package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInstructions(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func getValue(instructions []string) int {
	output := 0
	visited := make(map[int]bool)
	i := 0
	for {
		_, previouslySeen := visited[i]
		if previouslySeen {
			return output
		}
		visited[i] = true

		instruction := instructions[i]
		operation := instruction[:3]
		value, err := strconv.Atoi(strings.Split(instruction, " ")[1])
		if err != nil {
			panic(err)
		}

		if operation == "jmp" {
			i += value
			continue
		} else if operation == "acc" {
			output += value
		}
		i++
	}
}

func main() {
	instructions := readInstructions("input.txt")
	fmt.Println(getValue(instructions))
}
