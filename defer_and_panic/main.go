package main

import "fmt"

func main() {
	defer d1()
	fmt.Println(p(1, 22))
	fmt.Println("Starting")
	defer d2()
	fmt.Println(p(-1, 22))
	fmt.Println("Ending")
}

func d1() {
	fmt.Println("d1")
}

func d2() {
	fmt.Println("d2")
}

func p(x int, y int) int {
	if x < 0 {
		panic("x must be positivee ")
	} else {
		return x + y
	}
}
