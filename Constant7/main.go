package main

import "fmt"

func main() {
	const password = "123rasulQq"
	fmt.Println(password)
	const (
		pi float64 = 3.1415
		e  float64 = 2.7182
	)
	fmt.Println(pi, e)
	//var m int = 7
	//const k = m      // ! Ошибка: m - переменная
	const s = 5 // Норм: 5 - числовая константа
	const n = s // Норм: s - константа
}
