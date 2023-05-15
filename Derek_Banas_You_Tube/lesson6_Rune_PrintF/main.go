package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	rStr := "asdfghijklmnopqrstuv"
	fmt.Println("Rune Count:", utf8.RuneCountInString(rStr))
	for i, RuneV := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, RuneV, RuneV)
	}
}
