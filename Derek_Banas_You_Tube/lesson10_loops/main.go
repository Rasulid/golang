package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNumber := rand.Intn(10) + 1
	fmt.Println("Find the number in deoposone 0 to 10")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if num < randNumber {
			fmt.Println("The number is too low")
		} else if num > randNumber {
			fmt.Println("The number is too high")
		} else if num == randNumber {
			fmt.Println("The number is the same")
			fmt.Println("You win")
			fmt.Println("The number is", randNumber)
			break
		}

	}
}
