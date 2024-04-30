package main

import "fmt"

func factorial(n int) int {
	//factorial function calculates the factorial of using n recursion
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var input int
	fmt.Println("enter the Factorial number: ")
	fmt.Scanln(&input)
	result := factorial(input)
	fmt.Println("Factorial of", input, "is:", result)
}
