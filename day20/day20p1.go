package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readTiles(file string) map[int][]string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	output := make(map[int][]string)
	rawTiles := strings.Split(string(data), "\n\n")
	for _, rawTile := range rawTiles {
		lines := strings.Split(rawTile, "\n")
		var number int
		n, err := fmt.Sscanf(lines[0], "Tile %d:", &number)
		if n != 1 || err != nil {
			panic(err)
		}
		output[number] = lines[1:]
	}

	return output
}

func findCorners(tiles map[int][]string) int {
	// Maps a particular edge onto a set of tiles
	edgesMap := make(map[string][]int)
	for tileNumber, tile := range tiles {
		edges := extractEdges(tile)
		for _, edge := range edges {
			edgesMap[edge] = append(edgesMap[edge], tileNumber)
		}
	}

	// Maps number of unique edges per tile
	uniqueEdges := make(map[int]int)
	for _, tileNumbers := range edgesMap {
		if len(tileNumbers) == 1 {
			uniqueEdges[tileNumbers[0]]++
		}
	}

	multiple := 1
	for k, v := range uniqueEdges {
		if v == 4 {
			multiple *= k
		}
	}
	fmt.Println(uniqueEdges)
	return multiple
}

func extractEdges(tile []string) []string {
	set := make(map[string]bool)
	for _, s := range []string{tile[0], tile[len(tile)-1]} {
		set[s] = true
		set[reverseString(s)] = true
	}

	leftEdge := ""
	rightEdge := ""
	for _, line := range tile {
		leftEdge += string(line[0])
		rightEdge += string(line[len(line)-1])
	}
	for _, s := range []string{leftEdge, rightEdge} {
		set[s] = true
		set[reverseString(s)] = true
	}

	output := make([]string, 0)
	for s := range set {
		output = append(output, s)
	}
	return output
}

func reverseString(input string) string {
	output := ""
	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}
	return output
}

func main() {
	tiles := readTiles("input.txt")
	fmt.Println(tiles)
	fmt.Println(findCorners(tiles))
}
