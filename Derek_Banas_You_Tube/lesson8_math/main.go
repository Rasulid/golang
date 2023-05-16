package main

import (
	"fmt"
	"math"
)

func main() {
	//random := rand.Intn(50) + 1
	//println(random)
	fmt.Println("Abs(-10) =", math.Abs(-10))
	fmt.Println("Pow(4, 2) =", math.Pow(4, 2))
	fmt.Println("Sqrt(16) =", math.Sqrt(16))
	fmt.Println("Cbrt(8) =", math.Cbrt(8))
	fmt.Println("Ceil(4.4) =", math.Ceil(4.4))
	fmt.Println("Floor(4.4) =", math.Floor(4.4))
	fmt.Println("Round(4.4) =", math.Round(4.4))
	fmt.Println("Log2(8) =", math.Log2(8))
	fmt.Println("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	fmt.Println("Log(7.389) =", math.Log(math.Exp(2)))
	fmt.Println("Max(5,4) =", math.Max(5, 4))
	fmt.Println("Min(5,4) =", math.Min(5, 4))

	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians and %.2f degrees\n", r90, d90)

	fmt.Println("Sinus = ", math.Sin(r90))
}
