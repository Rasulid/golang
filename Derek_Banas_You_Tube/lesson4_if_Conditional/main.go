package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		Conditional Operators : > , <, <=, >= , == , !=
		Logical Operators : &&, ||, !
	*/

	fmt.Println("How old are you?")
	reader := bufio.NewReader(os.Stdin)
	_age, _ := reader.ReadString('\n')

	_age = strings.TrimSuffix(_age, "\n")
	_age = strings.TrimSuffix(_age, "\r")

	age, err := strconv.Atoi(_age)
	if err != nil {
		fmt.Println("Error :", err)
	}

	if age >= 1 && age <= 18 {
		fmt.Println("Your birthday is sooo important")
	} else {
		fmt.Println("Your birthday is not important\nGo to work bitch")
	}

}
