package main

import "fmt"

func main() {
	/*
			Varadic функция в языке Go (Golang) - это функция,
		которая может принимать переменное количество аргументов одного и того же типа.
		В официальной документации такие функции описываются как "variadic functions".
	*/

	fmt.Println(getSum(1, 2, 3, 4, 4, 5, 90))

}

func getSum(num ...int) int {
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum
}
