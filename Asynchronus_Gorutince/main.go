package main

import "fmt"

func main() {
	ch := make(chan int)
	//ch2 := make(chan string)
	go say("Hello world", ch)
	//time.Sleep(2 * time.Second)
	//fmt.Println(<-ch, <-ch2)
	for a := range ch {
		fmt.Println(a)
	}
	fmt.Println(<-ch)
}
func say(word string, ch chan int) {
	//fmt.Println(word)
	//ch <- 9
	//ch2 <- word

	for i := 0; i < 10; i++ {
		//fmt.Println(i)
		ch <- i
	}
	close(ch)
}
