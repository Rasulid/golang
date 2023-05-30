package main

import (
	"fmt"
	"strings"
)

func ReplaceDots(str string) string {
	return strings.Replace(str, ".", "-", 8)
}

func main() {
	fmt.Println(ReplaceDots("one.two.three"))
}
