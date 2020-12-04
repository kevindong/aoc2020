package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func readFile(file string) [][]string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("Couldn't read file")
	}

	fileContents := string(data)
	batched := strings.Split(fileContents, "\n\n")

	var output [][]string
	for _, passport := range batched {
		normalized := strings.ReplaceAll(passport, "\n", " ")
		splitted := strings.Split(normalized, " ")
		sort.Strings(splitted)
		output = append(output, splitted)
	}

	return output
}

func isValid(passport []string) bool {
	if len(passport) < 7 {
		return false
	}

	passportMap := make(map[string]string)
	for _, item := range passport {
		components := strings.Split(item, ":")
		passportMap[components[0]] = components[1]
	}

	expectedKeys := []string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}
	for _, expectedKey := range expectedKeys {
		_, found := passportMap[expectedKey]
		if !found {
			return false
		}
	}

	return true
}

func main() {
	passports := readFile("input.txt")
	valid := 0
	for _, passport := range passports {
		if isValid(passport) {
			fmt.Printf("%v is valid\n", passport)
			valid++
		} else {
			fmt.Printf("%v is invalid\n", passport)
		}
	}
	fmt.Println(valid)
}
