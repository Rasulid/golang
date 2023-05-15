package main

import "fmt"

func main() {

	var a = 4 // +
	var b = 6
	var c = a + b
	fmt.Println(c) // 10

	a = 4 // -
	b = 6
	c = a - b
	fmt.Println(c) // -2

	a = 4 // *
	b = 6
	c = a * b // 24
	fmt.Println(c)

	a = 10
	b = 4
	c = a / b
	fmt.Println(c) // 2

	var k float32 = 10
	var l float32 = 4
	var m float32 = k / l
	fmt.Println(m) // 2.5

	m = 10 / 4 // 2
	fmt.Println(m)

	m = 10 / 4.0 // 2.5
	fmt.Println(m)

	c = 35 % 3 // 2 (35 - 33 = 2)
	fmt.Println(c)

	a = 8
	a++
	fmt.Println(a) // 9

	a = 8
	a--
	fmt.Println(a)

}
