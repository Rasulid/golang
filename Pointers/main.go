package main

import "fmt"

func change(s *string) {
	*s = "LoL"
}

func main() {
	a := "LLL"
	fmt.Println(a)
	change(&a)
	fmt.Println(a)

}
