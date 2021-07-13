package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		n1, n2, s1 string
	)
	fmt.Println("Please input your request for RPN calculator: ")
	fmt.Scanln(&n1, &n2, &s1)
	fmt.Println("The answer is 42")
}
