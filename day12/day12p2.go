package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readMovements(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func processMovements(movements []string) int {
	waypointX := 10
	waypointY := 1

	currentX := 0
	currentY := 0

	for _, movement := range movements {
		command := movement[0]
		number, err := strconv.Atoi(movement[1:])
		if err != nil {
			panic("Couldn't convert number portion of movement")
		}

		if command == 'N' {
			waypointY += number
		} else if command == 'S' {
			waypointY -= number
		} else if command == 'E' {
			waypointX += number
		} else if command == 'W' {
			waypointX -= number
		} else if command == 'F' {
			for i := 0; i < number; i++ {
				currentX += waypointX
				currentY += waypointY
			}
		} else if command == 'L' || command == 'R' {
			mapping := map[string]int{
				"L90":  3,
				"L180": 2,
				"L270": 1,
				"R90":  1,
				"R180": 2,
				"R270": 3,
			}
			iterations, found := mapping[movement]
			if !found {
				panic("Unrecognized movement value")
			}
			for i := 0; i < iterations; i++ {
				originalX := waypointX
				originalY := waypointY
				waypointX = originalY
				waypointY = -originalX
			}
		} else {
			panic("Unrecognized command")
		}
	}

	return abs(currentX) + abs(currentY)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	movements := readMovements("input.txt")
	fmt.Println(processMovements(movements))
}
