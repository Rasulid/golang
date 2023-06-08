package main

import "fmt"

func main() {
	intSum := func(x, y int) int { return x + y }
	fmt.Printf("%d + %d = %d\n", 5, 6, intSum(5, 6))

	samPl1 := 1
	cHangeVar := func() {
		samPl1 += 1
	}
	cHangeVar()
	fmt.Println(samPl1)

	useFunction(SumVals, 6, 5)
}

func useFunction(f func(int, int) int, x, y int) {
	fmt.Println("function", f(x, y))

}

func SumVals(x, y int) int {
	return x + y
}
