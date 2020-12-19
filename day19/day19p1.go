package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func readInput(file string) ([]string, []string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	components := strings.Split(string(data), "\n\n")
	rules := strings.Split(components[0], "\n")
	for i := range rules {
		rules[i] = replaceAll(rules[i], " ", "  ")
		rules[i] += " "
	}
	return rules, strings.Split(components[1], "\n")
}

func generateRegex(rules []string) string {
	mapping := make(map[int]string)
	aIndex := ""
	bIndex := ""
	for _, rule := range rules {
		ruleNumbering := strings.Split(rule, ":")[0]
		if strings.Contains(rule, "a") {
			aIndex = " (" + ruleNumbering + ") "
		} else if strings.Contains(rule, "b") {
			bIndex = " (" + ruleNumbering + ") "
		}
	}
	for _, rule := range rules {
		if strings.HasSuffix(rule, `"a" `) || strings.HasSuffix(rule, `"b" `) {
			continue
		}
		ruleNumbering := toInt(strings.Split(rule, ":")[0])
		tentativeString := strings.Split(rule, ":")[1]
		tentativeString = replaceAll(tentativeString, aIndex, " (a) ")
		tentativeString = replaceAll(tentativeString, bIndex, " (b) ")
		mapping[ruleNumbering] = tentativeString
	}
	x := 0
	y := 1
	for len(mapping) > 1 {
		x = y
		y = len(mapping)
		if x == y {
			for k, v := range mapping {
				fmt.Printf("[%v]: [%v]\n\n", k, v)
			}
			panic("problem")
		}
		fmt.Println(len(mapping))
		temp := copyMap(mapping)
		for ruleNumbering, currentRegex := range mapping {
			if containsNumbers(currentRegex) {
				continue
			}
			delete(temp, ruleNumbering)
			for k, v := range temp {
				temp[k] = replaceAll(v, " "+strconv.Itoa(ruleNumbering)+" ", " ("+currentRegex+") ")
				temp[k] = replaceAll(temp[k], "   ", "  ")
			}
		}
		mapping = temp
	}

	return replaceAll(mapping[0], " ", "")
}

func containsNumbers(input string) bool {
	matched, err := regexp.MatchString(`\d`, input)
	if err != nil {
		panic(err)
	}
	return matched
}

func replaceAll(source, find, replacement string) string {
	re := regexp.MustCompile(find)
	return re.ReplaceAllString(source, replacement)
}

func copyMap(original map[int]string) map[int]string {
	copy := make(map[int]string)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

func toInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	inputFile := "input.txt"
	fmt.Println("Running on:", inputFile)
	rules, testStrings := readInput(inputFile)
	regex := "^(" + generateRegex(rules) + ")$"
	fmt.Println(regex)

	matchCount := 0
	for _, testString := range testStrings {
		matched, err := regexp.MatchString(regex, testString)
		if err != nil {
			panic(err)
		}
		if matched {
			fmt.Println(testString, "matches")
			matchCount++
		} else {
			fmt.Println(testString, "does not match")
		}
	}
	fmt.Println(matchCount)
}
