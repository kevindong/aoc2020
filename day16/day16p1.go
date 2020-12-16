package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(file string) (map[int]bool, []string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	components := strings.Split(string(data), "\n\n")
	validNumbers := make(map[int]bool)

	rulesString := strings.Split(components[0], "\n")
	for _, ruleLine := range rulesString {
		ruleComponents := strings.Split(ruleLine, " ")
		for _, ruleComponent := range ruleComponents {
			if strings.Contains(ruleComponent, "-") {
				ranges := strings.Split(ruleComponent, "-")
				start, err := strconv.Atoi(ranges[0])
				if err != nil {
					panic(err)
				}
				end, err := strconv.Atoi(ranges[1])
				if err != nil {
					panic(err)
				}
				for i := start; i <= end; i++ {
					validNumbers[i] = true
				}
			}
		}
	}

	tickets := strings.Split(components[2], "\n")

	return validNumbers, tickets[1:]
}

func checkBasicTicketValidity(validNumbers map[int]bool, tickets []string) int {
	counter := 0
	for _, ticket := range tickets {
		ticketValues := strings.Split(ticket, ",")
		for _, ticketValue := range ticketValues {
			number, err := strconv.Atoi(ticketValue)
			if err != nil {
				panic(err)
			}

			if !validNumbers[number] {
				counter += number
			}
		}
	}
	return counter
}

func main() {
	validNumbers, tickets := readInput("input.txt")
	//fmt.Println(validNumbers)
	//fmt.Println(tickets)
	fmt.Println(checkBasicTicketValidity(validNumbers, tickets))
}
