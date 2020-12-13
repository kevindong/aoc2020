package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(file string) (int, []int) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	departureTimestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}

	busIdsString := strings.Split(lines[1], ",")
	busIdsInts := make([]int, 0)
	for _, busId := range busIdsString {
		id, err := strconv.Atoi(busId)
		if err != nil {
			id = -1
		}
		busIdsInts = append(busIdsInts, id)
	}

	return departureTimestamp, busIdsInts
}

func main() {
	fmt.Println("Use a Chinese remainder theorem with these inputs:")
	_, buses := readInput("input.txt")
	for i, busId := range buses {
		if busId != -1 {
			fmt.Printf("Remainder %v with mod %v\n", i, busId)
		}
	}
}
