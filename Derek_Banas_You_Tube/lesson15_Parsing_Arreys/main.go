package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4}
	fmt.Println("Sum of arrays:", getSumArray(array))
}

func getSumArray(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}
