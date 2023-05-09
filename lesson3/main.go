package main

import (
	"fmt"
	"math"
)

func main() {
	//num := 3
	//if num > 3 {
	//	fmt.Println("Number is greater then ", num)
	//} else if num < 3 {
	//	fmt.Println("Number is lesser then ", num)
	//} else if num == 3 {
	//	fmt.Println("Number is equals to ", num)
	//} else {
	//	fmt.Println("working else ", num)
	//}

	var a float64
	var b float64
	var c float64

	fmt.Println("RESOLVE MATH")

	fmt.Println("Enter a number: ")
	fmt.Scan(&a)

	fmt.Println("Enter b number: ")
	fmt.Scan(&b)

	fmt.Println("Enter c number: ")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уравнения имеет 2 корня \nD = " + fmt.Sprint(D))
		fmt.Println("x1 = "+fmt.Sprint(x1), "x2 = "+fmt.Sprint(x2))
	} else if D == 0 {
		var x1 float64
		fmt.Println("Ваше уравнения имеет 0 корня \nD = " + fmt.Sprint(D))
		fmt.Println("x = " + fmt.Sprint(x1))
	} else if D < 0 {
		fmt.Println("Ваше уравнения не имеет корней ")
		fmt.Println("D = " + fmt.Sprint(D))
	}

}
