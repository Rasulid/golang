package main

import "fmt"

func main() {
	f3 := 10
	f4 := 20
	var f2Pint *int = &f4
	fmt.Println("F4 address: ", f2Pint)
	fmt.Println("F4 value", *f2Pint)
	*f2Pint = 11
	fmt.Println("F4 value", *f2Pint)
	fmt.Println("F3 before function", f3)
	changeVal(&f3)
	fmt.Println("F3 after function", f3)
}

func changeVal(myPtr *int) {
	*myPtr = 12
}
