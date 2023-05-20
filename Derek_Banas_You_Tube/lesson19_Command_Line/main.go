package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	args := os.Args[1:]
	var iArgs = []int{}
	for _, a := range args {
		val, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	max := 0
	for _, a := range iArgs {
		if a > max {
			max = a
		}
	}
	fmt.Printf("%d is max value\n", max)
}
