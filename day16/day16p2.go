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

func checkBasicTicketValidity(validNumbers map[int]bool, tickets []string) [][]int {
	validTickets := make([][]int, 0)
	for _, ticket := range tickets {
		ticketValues := strings.Split(ticket, ",")
		intTicket := make([]int, 0)
		isValid := true
		for _, ticketValue := range ticketValues {
			number, err := strconv.Atoi(ticketValue)
			if err != nil {
				panic(err)
			}
			intTicket = append(intTicket, number)

			if !validNumbers[number] {
				isValid = false
				break
			}
		}
		if isValid {
			validTickets = append(validTickets, intTicket)
		}
	}
	return validTickets
}

func mapValuesToRules(file string) map[int][]string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	mapping := make(map[int][]string)
	rules := strings.Split(strings.Split(string(data), "\n\n")[0], "\n")
	for _, rule := range rules {
		field := strings.Split(rule, ":")[0]
		ranges := strings.Split(strings.Split(rule, ": ")[1], " or ")
		for _, r := range ranges {
			numbers := strings.Split(r, "-")
			start, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}

			for i := start; i <= end; i++ {
				mapping[i] = append(mapping[i], field)
			}
		}
	}
	return mapping
}

func findMissing(fieldValue int, mapping map[int][]string) string {
	var completeFieldList []string
	for _, list := range mapping {
		if len(list) == 20 {
			completeFieldList = list
			break
		}
	}

	have := make(map[string]bool)
	for _, field := range mapping[fieldValue] {
		have[field] = true
	}

	for _, field := range completeFieldList {
		if !have[field] {
			return field
		}
	}

	return ""
}

func getAllFields(mapping map[int][]string) map[string]bool {
	var allFieldsList []string
	for _, list := range mapping {
		if len(list) == 20 {
			allFieldsList = list
		}
	}
	allFieldMap := make(map[string]bool)
	for _, field := range allFieldsList {
		allFieldMap[field] = true
	}
	return allFieldMap
}

func main() {
	inputFile := "input.txt"
	validNumbers, tickets := readInput(inputFile)
	validTickets := checkBasicTicketValidity(validNumbers, tickets)
	mapping := mapValuesToRules(inputFile)
	counter := make(map[int]int)

	validMappings := make(map[int]map[string]bool)
	for i := range validTickets[0] {
		validMappings[i] = getAllFields(mapping)
	}
	for _, ticket := range validTickets {
		for i, fieldValue := range ticket {
			if mapping[fieldValue] == nil {
				panic("Supposedly invalid field found")
			}

			counter[len(mapping[fieldValue])]++
			if len(mapping[fieldValue]) != 20 {
				missingType := findMissing(fieldValue, mapping)
				delete(validMappings[i], missingType)
			}
		}
	}

	fields := getAllFields(mapping)
	for i := 1; i <= 20; i++ {
		for fieldIndex, validFields := range validMappings {
			if len(validFields) == i {
				for field := range validFields {
					if fields[field] {
						fields[field] = false
						fmt.Printf("Field index %v is of type %v\n", fieldIndex, field)
						break
					}
				}
				break
			}
		}
	}

	fmt.Println("I'm just over this problem, but at this point you have a 0-indexed list of which columns correspond to which fields so it should just be some simple math.")
}
