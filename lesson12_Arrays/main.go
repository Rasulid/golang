package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arre := [...]int{1, 2, 3, 4, 5, 32, 54, 23}
	arre[5] = 4332
	//fmt.Println(arre[5])

	dictionary := []string{2: "Rasul", 94: "Abdi"}
	fmt.Println(dictionary[2])
}
