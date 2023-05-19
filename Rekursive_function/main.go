package main

import "fmt"

func main() {
	fmt.Println(recourse(6))
	fmt.Println(fibonacci(3))
}

func recourse(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * recourse(n-1)
}

func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
