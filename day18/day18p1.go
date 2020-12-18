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
	stack := make([]string, 0)
	for _, token := range input {
		if len(stack) == 0 {
			stack = append(stack, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			if len(stack) == 2 {
				stack = stack[1:]
			} else {
				// [... x _ ( 1 ]
				lastOperand := convertStringToInt(stack[len(stack)-1])
				operator := stack[len(stack)-3]
				if operator == "(" {
					stack = stack[:len(stack)-2]
					stack = append(stack, strconv.Itoa(lastOperand))
					continue
				}
				penultimateOperand := convertStringToInt(stack[len(stack)-4])
				result := 0
				if operator == "+" {
					result = lastOperand + penultimateOperand
				} else if operator == "*" {
					result = lastOperand * penultimateOperand
				} else {
					panic("unrecognized operand")
				}
				// TODO: This might not be right
				stack = stack[:len(stack)-4]
				stack = append(stack, strconv.Itoa(result))
			}
		} else if token == "+" || token == "*" {
			stack = append(stack, token)
		} else { // We're in a number
			number := convertStringToInt(token)
			lastToken := stack[len(stack)-1]
			if lastToken == "+" || lastToken == "*" {
				operator := lastToken
				operand := convertStringToInt(stack[len(stack)-2])
				result := 0
				if operator == "+" {
					result = number + operand
				} else if operator == "*" {
					result = number * operand
				} else {
					panic("unrecognized operand")
				}
				stack = stack[:len(stack)-2]
				stack = append(stack, strconv.Itoa(result))
			} else if lastToken == "(" {
				stack = append(stack, token)
			} else {
				panic("unexpected token")
			}
		}
	}
	return convertStringToInt(stack[0])
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
		sum += doProblem(problem)
	}
	fmt.Println(sum)
}
