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
			continue
		}
		busIdsInts = append(busIdsInts, id)
	}

	return departureTimestamp, busIdsInts
}

func getAnswer(departureTimestamp int, busIds []int) int {
	min := 100000
	selectedBusId := -1
	for _, busId := range busIds {
		waitTime := busId - (departureTimestamp % busId)
		if waitTime < min {
			min = waitTime
			selectedBusId = busId
		}
	}
	return min * selectedBusId
}

func main() {
	departureTimestamp, busIds := readInput("input.txt")
	fmt.Println(getAnswer(departureTimestamp, busIds))
}
