package main

import (
	"fmt"
	"math"
)

func main() {
	// %Method : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base8
	// %x : Base16
	// %v : Guasses based on data type
	// %T : Type of suplited value

	fmt.Printf("%s %Method %c %.2f %t %o %x\n",
		"Hello", 123, 'a', 3.14, true, 1, 1)

	fmt.Printf("%9f\n", 1.34)
	fmt.Printf("%.2f\n", 1.3423421434)
	fmt.Printf("%9.f\n", 3.341235464)

	sp1 := fmt.Sprintf("%9.f", math.Pi)
	println(sp1)

}
