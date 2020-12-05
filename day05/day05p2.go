package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readSeats(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("Couldn't read file.")
	}

	return strings.Split(string(data), "\n")
}

func decode(code string) int {
	left := 0
	var right int
	if len(code) == 7 {
		right = 127
	} else {
		right = 7
	}

	for i := 0; i < len(code); i++ {
		if code[i:i+1] == "F" {
			right = (left + right) / 2
		} else {
			left = ((left + right) / 2) + 1
		}
	}

	return left
}

func getSeatID(row, col int) int {
	return (row * 8) + col
}

func main() {
	seats := readSeats("input.txt")
	maxSeatID := 0
	foundSeats := make(map[int]bool)
	for _, seat := range seats {
		normalizedSeat := strings.ReplaceAll(seat, "L", "F")
		normalizedSeat = strings.ReplaceAll(normalizedSeat, "R", "B")
		row := decode(normalizedSeat[:7])
		col := decode(normalizedSeat[7:])
		seatID := getSeatID(row, col)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
		foundSeats[seatID] = true
	}
	fmt.Println(maxSeatID)
	for i := 0; i < maxSeatID; i++ {
		_, priorSeatFound := foundSeats[i-1]
		_, currentSeatFound := foundSeats[i]
		_, nextSeatFound := foundSeats[i+1]
		if priorSeatFound && !currentSeatFound && nextSeatFound {
			fmt.Println(i)
			return
		}
	}
}
