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
	currentX := 0
	currentY := 0
	currentDirection := "E"
	currentDegrees := 90

	for _, movement := range movements {
		fmt.Printf("(%v, %v) facing (%v = %v)\n", currentX, currentY, currentDirection, currentDegrees)
		command := string(movement[0])
		number, err := strconv.Atoi(movement[1:])
		if err != nil {
			panic(err)
		}
		fmt.Println("Parsed", command, number)

		if command == "N" {
			currentY += number
		} else if command == "S" {
			currentY -= number
		} else if command == "E" {
			currentX += number
		} else if command == "W" {
			currentX -= number
		} else if command == "F" {
			if currentDirection == "N" {
				currentY += number
			} else if currentDirection == "S" {
				currentY -= number
			} else if currentDirection == "E" {
				currentX += number
			} else if currentDirection == "W" {
				currentX -= number
			} else {
				panic("Unrecognized direction")
			}
		} else if command == "L" {
			currentDegrees += 360 - number
			currentDegrees = currentDegrees % 360
		} else if command == "R" {
			currentDegrees += number
			currentDegrees = currentDegrees % 360
		} else {
			panic("Unrecognized command")
		}

		if currentDegrees == 0 {
			currentDirection = "N"
		} else if currentDegrees == 90 {
			currentDirection = "E"
		} else if currentDegrees == 180 {
			currentDirection = "S"
		} else if currentDegrees == 270 {
			currentDirection = "W"
		} else {
			panic("Unknown degree")
		}
	}

	fmt.Printf("(%v, %v) facing (%v = %v)\n", currentX, currentY, currentDirection, currentDegrees)
	return abs(currentX) + abs(currentY)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	movements := readMovements("input.txt")
	fmt.Println(processMovements(movements))
}
