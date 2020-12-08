package main

import (
	"errors"
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

func getValue(instructions []string) (int, error) {
	output := 0
	visited := make(map[int]bool)
	i := 0
	for {
		if visited[i] {
			return 0, errors.New("repeated instruction")
		} else if i >= len(instructions) {
			return output, nil
		}
		visited[i] = true

		instruction := instructions[i]
		operation := instruction[:3]
		value, err := strconv.Atoi(strings.Split(instruction, " ")[1])
		if err != nil {
			panic(err)
		}

		if operation == "nop" {
			i++
		} else if operation == "jmp" {
			i += value
		} else if operation == "acc" {
			output += value
			i++
		}
	}
}

func main() {
	instructions := readInstructions("input.txt")
	for i, instruction := range instructions {
		original := instruction
		if original[:3] == "acc" || original == "nop +0" {
			continue
		}

		operation := original[:3]
		if operation == "nop" {
			operation = "jmp"
		} else {
			operation = "nop"
		}

		instructions[i] = operation + original[3:]
		output, err := getValue(instructions)
		if err == nil {
			fmt.Println(output)
			return
		}
		instructions[i] = original
	}
}
