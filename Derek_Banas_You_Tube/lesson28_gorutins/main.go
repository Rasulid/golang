package main

import (
	"fmt"
	"time"
)

func PrintTo10() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Function 1 value %d\n", i)
	}
}

func PrintTo15() {
	for i := 0; i < 15; i++ {
		fmt.Printf("Function 2 value %d\n", i)
	}
}

func main() {
	go PrintTo10()
	go PrintTo15()

	time.Sleep(2 * time.Second)
}
