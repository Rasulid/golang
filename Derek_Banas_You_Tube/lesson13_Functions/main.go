package main

import "fmt"

func main() {
	fmt.Println(getSum(5, 10))
	fmt.Println(getQuotient(10, 0))
}

func getSum(x int, y int) int {
	println("Working")
	return x + y
}

func getQuotient(x float64, y float64) (aswer float64,
	err error) {
	if y == 0 {
		return 0, fmt.Errorf("you cann not devide by zero")
	} else {
		return x / y, nil
	}

}
