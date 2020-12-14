package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(file string) [][]string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(data), "mask = ")
	output := make([][]string, 0)
	for _, group := range s {
		if group == "" {
			continue
		}
		temp := strings.Split(group, "\n")
		output = append(output, temp)
	}
	return output
}

func maskAddress(bitmask string, address string) []int64 {
	n, err := strconv.Atoi(address)
	if err != nil {
		panic(err)
	}
	addressInt := int64(n)
	addressBitString := strconv.FormatInt(addressInt, 2)
	addressBitString = fmt.Sprintf("%036s", addressBitString)

	outputAddress := ""
	for i, bit := range bitmask {
		if bit == '0' {
			outputAddress += string(addressBitString[i])
		} else if bit == '1' {
			outputAddress += "1"
		} else if bit == 'X' {
			outputAddress += "X"
		} else {
			panic("Unrecognized bit")
		}
	}

	return getAddresses(outputAddress)
}

func getAddresses(bitmask string) []int64 {
	workingSet := []string{""}
	for _, bit := range bitmask {
		if bit != 'X' {
			for s := range workingSet {
				workingSet[s] += string(bit)
			}
			continue
		}

		temp := make([]string, 0)
		for _, s := range workingSet {
			temp = append(temp, s+"0")
			temp = append(temp, s+"1")
		}
		workingSet = temp
	}

	output := make([]int64, 0)
	for _, address := range workingSet {
		convertedInt, err := strconv.ParseInt(address, 2, 64)
		if err != nil {
			panic(err)
		}
		output = append(output, convertedInt)
	}
	return output
}

func solve(group []string, mapping map[int64]int64) {
	mask := group[0]
	for i := 1; i < len(group); i++ {
		rule := group[i]
		if rule == "" {
			continue
		}
		locationString := strings.TrimPrefix(strings.Split(rule, "]")[0], "mem[")
		addresses := maskAddress(mask, locationString)

		numberString := strings.Split(rule, " ")[2]
		n, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}
		num := int64(n)
		for _, address := range addresses {
			mapping[address] = num
		}
	}
}

func main() {
	mapping := make(map[int64]int64)
	groups := readInput("input.txt")
	for _, group := range groups {
		solve(group, mapping)
	}

	sum := int64(0)
	for _, value := range mapping {
		sum += value
	}
	fmt.Println(sum)
}
