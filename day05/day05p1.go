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

// code should be normalized to be F and B only.
// returns the 0-indexed value
func decode(code string) int {
	left := 1
	var right int
	if len(code) == 7 {
		right = 128
	} else {
		right = 8
	}

	for i := 0; i < len(code); i++ {
		if code[i:i+1] == "F" {
			right = (left + right) / 2
			//fmt.Println("Right is now ", right)
		} else {
			left = (left + right) / 2
			//fmt.Printf("HAH: (%v - %v + 1) / 2\n", right, left)
			//fmt.Println("Left is now ", left)
		}
	}

	//fmt.Println("Left is", left, "; Right is", right, "\n")

	return left
}

func getSeatID(row, col int) int {
	return (row * 8) + col
}

func main() {
	seats := readSeats("input.txt")
	max := -1
	for _, seat := range seats {
		normalizedSeat := strings.ReplaceAll(seat, "L", "F")
		normalizedSeat = strings.ReplaceAll(normalizedSeat, "R", "B")
		row := decode(normalizedSeat[:7])
		col := decode(normalizedSeat[7:])
		seatID := getSeatID(row, col)
		//fmt.Printf("Row = %v; col = %v\n", normalizedSeat[:7], normalizedSeat[7:])
		fmt.Printf("Seat (%v) returns row (%v) and col (%v) with seatID %v\n\n", seat, row, col, seatID)
		if seatID > max {
			max = seatID
		}
	}
	fmt.Println(max)
}
