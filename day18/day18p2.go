package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readProblems(file string) [][]string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	problems := strings.Split(string(data), "\n")
	output := make([][]string, 0)
	for _, problem := range problems {
		tokens := make([]string, 0)
		for _, r := range problem {
			if r != ' ' {
				tokens = append(tokens, string(r))
			}
		}
		output = append(output, tokens)
	}
	return output
}

func doProblem(input []string) int {
	operatorStack := make([]string, 0)
	numberStack := make([]int, 0)
	for _, token := range input {
		if token == "(" {
			operatorStack = append(operatorStack, token)
		} else if token == ")" {
			factors := make([]int, 0)
			factors = append(factors, numberStack[len(numberStack)-1])
			numberStack = numberStack[:len(numberStack)-1]
			for operatorStack[len(operatorStack)-1] == "*" {
				factors = append(factors, numberStack[len(numberStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
				numberStack = numberStack[:len(numberStack)-1]
			}
			if operatorStack[len(operatorStack)-1] == "(" {
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			result := factors[0]
			for _, factor := range factors[1:] {
				result *= factor
			}
			numberStack = append(numberStack, result)
			if len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] == "+" {
				operatorStack = operatorStack[:len(operatorStack)-1]
				lastNumber := numberStack[len(numberStack)-1]
				numberStack = numberStack[:len(numberStack)-1]
				numberStack[len(numberStack)-1] += lastNumber
			}
		} else if token == "+" || token == "*" {
			operatorStack = append(operatorStack, token)
		} else { // It's a number
			number := convertStringToInt(token)
			if len(operatorStack) == 0 || operatorStack[len(operatorStack)-1] == "(" {
				numberStack = append(numberStack, number)
				continue
			}
			lastOperator := operatorStack[len(operatorStack)-1]
			if lastOperator == "+" {
				operatorStack = operatorStack[:len(operatorStack)-1]
				numberStack[len(numberStack)-1] += number
			} else if lastOperator == "*" {
				numberStack = append(numberStack, number)
			}
		}
	}

	result := numberStack[0]
	for _, number := range numberStack[1:] {
		result *= number
	}
	return result
}

func convertStringToInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	input := "input.txt"
	fmt.Println("Running on", input)
	problems := readProblems(input)
	sum := 0
	for _, problem := range problems {
		result := doProblem(problem)
		fmt.Println(problem, "=", result)
		sum += result
	}
	fmt.Println(sum)
}
