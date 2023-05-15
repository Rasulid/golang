package main

import "fmt"

func main() {
	x := 100
	y := 111
	c := x ^ y
	fmt.Println(c)
	var d int = 5 &^ 6
	fmt.Println(d)

	c = 16 >> 3
	var b int = 2 << 2
	fmt.Println(c, b)
}
