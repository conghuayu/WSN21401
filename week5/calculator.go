package main

import (
	"fmt"
	"strconv"
)

func calRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {

		if val, err := strconv.Atoi(v); err == nil {
			stack = append(stack, val)
		} else {
			if len(stack) < 2 {
				return 0
			}
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			fmt.Println(n1, n2, v)
			stack = stack[:len(stack)-2]
			n := 0
			switch v {
			case "+":
				n += n1 + n2
			case "-":
				n += n1 - n2
			case "*":
				n += n1 * n2
			case "/":
				n += n1 / n2
			default:
				return 0
			}
			stack = append(stack, n)
		}
	}
	return stack[len(stack)-1]
}

func main() {
	var (
		n1, n2, s1 string
	)
	fmt.Println("Please input your request for RPN calculator: ")
	fmt.Scanln(&n1, &n2, &s1)
	n := []string{n1, n2, s1}
	result := calRPN(n)
	//fmt.Println("The answer is 42")
	fmt.Println("The answer is " + strconv.Itoa(result))
}
