package main

import "fmt"

func main() {
	//	var age int8 = 28
	//	fmt.Println(age)
	//	var f float32 = 12.3
	//	fmt.Println(f)
	//	var s string = "Hello my name is Rasul"
	//	fmt.Println(s)
	//	var uns uint8 = 123 // unsight integer only positive number
	//	fmt.Println(uns)

	//age := 12 // without short writing the variables
	//fmt.Println(age)

	//name := "Rasul"
	//fmt.Println(name)

	var name string
	fmt.Println("what is your name")
	fmt.Scan(&name)
	fmt.Println("hello " + name + " !")
	var age int16
	fmt.Println("how old are you")
	fmt.Scan(&age)
	fmt.Println("nice "+name+" you are ", age)
	fmt.Println(len(name))
	fmt.Println("you are " + fmt.Sprint(age) + " old")

}
