package main

import (
	"fmt"
)

func main() {
	var (
		n1, n2, s1 string
	) //n-number,s-symbol
	fmt.Println("Please input your request for RPN calculator: ")
	fmt.Scanln(&n1, &n2, &s1)
	n := []string{n1, n2, s1}
	//result := calRPN(n)
	fmt.Println("The answer is 42")
}
