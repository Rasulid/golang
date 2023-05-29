package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}
type NumbersInterface interface {
	Sum() int
	Multiply() int
	Divide() float64
	Subtract() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}
func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}
func (n Numbers) Divide() float64 {
	return float64(n.num1) / float64(n.num2)
}
func (n Numbers) Subtract() int {
	return n.num1 - n.num2
}

func main() {
	var i NumbersInterface
	Number := Numbers{9, 9}
	i = Number
	fmt.Println(i.Divide())

}
