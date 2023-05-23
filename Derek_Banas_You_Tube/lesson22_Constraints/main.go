package main

import "fmt"

type MyConstraints interface {
	int | float64
}

func main() {
	fmt.Println(getSum(9, 6))
	fmt.Println(getSum(9.9, 6.6))
}

func getSum[T MyConstraints](x T, y T) T {
	return x * y
}
