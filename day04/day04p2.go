package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
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

	birthYearRaw, found := passportMap["byr"]
	if !found || len(birthYearRaw) != 4 {
		return false
	}
	birthYear, err := strconv.Atoi(birthYearRaw)
	if err != nil || birthYear < 1920 || birthYear > 2002 {
		return false
	}

	issueYearRaw, found := passportMap["iyr"]
	if !found || len(issueYearRaw) != 4 {
		return false
	}
	issueYear, err := strconv.Atoi(issueYearRaw)
	if err != nil || issueYear < 2010 || issueYear > 2020 {
		return false
	}

	expirationYearRaw, found := passportMap["eyr"]
	if !found || len(expirationYearRaw) != 4 {
		return false
	}
	expirationYear, err := strconv.Atoi(expirationYearRaw)
	if err != nil || expirationYear < 2020 || expirationYear > 2030 {
		return false
	}

	heightRaw, found := passportMap["hgt"]
	if !found || len(heightRaw) < 4 {
		return false
	}
	units := heightRaw[len(heightRaw)-2:]
	if units != "cm" && units != "in" {
		return false
	}
	heightNumberRaw := heightRaw[:len(heightRaw)-2]
	heightNumber, err := strconv.Atoi(heightNumberRaw)
	fmt.Printf("Parsed height (%v) to %v %v\n", heightRaw, heightNumber, units)
	if err != nil {
		return false
	}
	if units == "cm" && (heightNumber < 150 || heightNumber > 193) {
		return false
	} else if units == "in" && (heightNumber < 59 || heightNumber > 76) {
		return false
	}

	hairColor, found := passportMap["hcl"]
	if !found {
		return false
	}
	matched, err := regexp.MatchString(`#[A-z0-9]{6}`, hairColor)
	if err != nil || !matched {
		return false
	}

	eyeColor, found := passportMap["ecl"]
	if !found {
		return false
	}
	eyeColorIsValid := false
	validColors := strings.Split("amb blu brn gry grn hzl oth", " ")
	for _, validColor := range validColors {
		if validColor == eyeColor {
			eyeColorIsValid = true
			break
		}
	}
	if !eyeColorIsValid {
		return false
	}

	passportID, found := passportMap["pid"]
	if !found || len(passportID) != 9 {
		return false
	}
	_, err = strconv.Atoi(passportID)
	if err != nil {
		return false
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
