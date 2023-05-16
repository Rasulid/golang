package main

import (
	"fmt"
	"time"
)

func main() {
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//
	//go func() {
	//	ch1 <- 1 // запись в канал ch1
	//}()
	//
	//select {
	//case <-ch1:
	//	fmt.Println("Прочитано из ch1")
	//case <-ch2:
	//	fmt.Println("Прочитано из ch2")
	//default:
	//	fmt.Println("Ни одно действие не может быть выполнено немедленно")
	//}

	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(6 * time.Second)
		one <- "one"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		two <- "two"
	}()

	select {
	case <-one:
		fmt.Println("one")
	case <-two:
		fmt.Println("two")
	default:
		fmt.Println("Default")

	}

}
