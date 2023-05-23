package main

import (
	stuff "exaple/project/mypackege"
	"fmt"
	"log"
	"reflect"
)

func main() {
	fmt.Println("hello", stuff.Name)
	ss := []int{1, 2, 3, 4}
	sF := stuff.IntArrToStrArr(ss)
	fmt.Println(sF)
	fmt.Println(reflect.TypeOf(sF))

	date := stuff.Date{}

	err := date.SetYear(2003)
	if err != nil {
		panic(err)
	}

	err = date.SetMonth(11)
	if err != nil {
		panic(err)
	}

	err = date.SetDay(30)
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}

	fmt.Printf("My birthday on %d/%d/%d",
		date.Day(), date.Month(), date.Year())
}
