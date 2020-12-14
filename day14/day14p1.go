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

func maskNumber(bitmask string, number string) int64 {
	n, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	num := int64(n)

	for i, bit := range bitmask {
		complement := int64(len(bitmask) - i - 1)
		if bit == '1' {
			num = setBitOne(num, complement)
		} else if bit == '0' {
			num = clearBit(num, complement)
		}
	}
	fmt.Println(bitmask, number, num)
	return num
}

func setBitOne(n int64, pos int64) int64 {
	n |= (1 << pos)
	return n
}

func clearBit(n int64, pos int64) int64 {
	mask := int64(^(1 << pos))
	n &= mask
	return n
}

func solve(group []string, mapping map[string]int64) {
	mask := group[0]
	for i := 1; i < len(group); i++ {
		rule := group[i]
		if rule == "" {
			continue
		}
		locationString := strings.Split(rule, " ")[0]
		numberString := strings.Split(rule, " ")[2]
		mapping[locationString] = maskNumber(mask, numberString)
	}
}

func main() {
	mapping := make(map[string]int64)
	groups := readInput("input.txt")
	for _, group := range groups {
		// fmt.Println(strings.Join(group, "\n"))
		solve(group, mapping)
	}

	sum := int64(0)
	for _, value := range mapping {
		sum += value
	}
	fmt.Println(sum)
}
