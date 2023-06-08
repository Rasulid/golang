package main

import "fmt"

func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	
	return num * Factorial(num-1)
}

func main() {
	fmt.Println(Factorial(5))
}
