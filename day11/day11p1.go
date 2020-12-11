package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(file string) [][]string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	output := make([][]string, 0)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		output = append(output, strings.Split(line, ""))
	}
	return output
}

func process(seatmap [][]string) ([][]string, int) {
	changed := 0

	temp := duplicate(seatmap)
	for row := range seatmap {
		for col := range seatmap[row] {
			if seatmap[row][col] == "." {
				continue
			}
			adjacentOccupiedSeats := countOccupiedAdjacentSeats(seatmap, row, col)
			if seatmap[row][col] == "L" && adjacentOccupiedSeats == 0 {
				temp[row][col] = "#"
				changed++
			} else if seatmap[row][col] == "#" && adjacentOccupiedSeats >= 4 {
				temp[row][col] = "L"
				changed++
			}
		}
	}

	return temp, changed
}

func duplicate(input [][]string) [][]string {
	output := make([][]string, 0)
	for _, row := range input {
		temp := make([]string, len(row))
		copy(temp, row)
		output = append(output, temp)
	}
	return output
}

func countOccupiedAdjacentSeats(seatmap [][]string, row, col int) int {
	surroundingDiffs := [][]int{
		{-1, -1}, // Top left
		{-1, 0},  // Top
		{-1, 1},  // Top right
		{0, -1},  // Left
		{0, 1},   // Right
		{1, -1},  // Bottom left
		{1, 0},   // Bottom
		{1, 1},   // Bottom right
	}

	occupied := 0
	for _, surroundingDiff := range surroundingDiffs {
		currentRow := row + surroundingDiff[0]
		currentCol := col + surroundingDiff[1]
		// Out of bounds; skip
		if currentRow < 0 || currentRow >= len(seatmap) || currentCol < 0 || currentCol >= len(seatmap[0]) {
			continue
		}

		if seatmap[currentRow][currentCol] == "#" {
			occupied++
		}
	}
	return occupied
}

func prettyPrint(input [][]string) {
	for _, row := range input {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println("")
}

func main() {
	lines := readFile("input.txt")
	for {
		changed := 0
		lines, changed = process(lines)
		if changed == 0 {
			break
		}
	}

	occupied := 0
	for _, row := range lines {
		for _, seat := range row {
			if seat == "#" {
				occupied++
			}
		}
	}
	fmt.Println(occupied)
}
