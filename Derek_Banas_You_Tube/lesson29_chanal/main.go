package main

import "fmt"

func num1(chen chan int) {
	chen <- 1
	chen <- 2
	chen <- 3
}

func num2(chen chan int) {
	chen <- 4
	chen <- 5
	chen <- 6
}

func main() {
	chanal1 := make(chan int)
	chanal2 := make(chan int)

	go num1(chanal1)
	go num2(chanal2)

	fmt.Println(<-chanal1)
	fmt.Println(<-chanal1)
	fmt.Println(<-chanal1)
	fmt.Println(<-chanal2)
	fmt.Println(<-chanal2)
	fmt.Println(<-chanal2)

}
