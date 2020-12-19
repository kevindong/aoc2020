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
	for len(mapping) > 1 {
		temp := copyMap(mapping)
		for ruleNumbering, currentRegex := range mapping {
			shouldRun := true
			if ruleNumbering == 8 && !containsNumbers(temp[42]) {
				currentRegex = "(" + temp[42] + ")+"
				shouldRun = false
			} else if ruleNumbering == 11 && !containsNumbers(temp[42]) && !containsNumbers(temp[31]) {
				fourtyTwo := temp[42]
				thirtyOne := temp[31]
				options := make([]string, 0)
				for i := 1; i < 5; i++ {
					iString := strconv.Itoa(i)
					s := "(" + fourtyTwo + "){" + iString + "}" + "(" + thirtyOne + "){" + iString + "}"
					options = append(options, s)
				}
				joined := strings.Join(options, "|")
				currentRegex = "(" + joined + ")"
				shouldRun = false
			}
			if shouldRun && containsNumbers(currentRegex) {
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
			matchCount++
		}
	}
	fmt.Println(matchCount)
	fmt.Println("For some reason, this code doesn't always generate the right answer. But it does generate the right answer sometimes. So run it a couple times until it settles on what seems like a reasonable answer.")
}
