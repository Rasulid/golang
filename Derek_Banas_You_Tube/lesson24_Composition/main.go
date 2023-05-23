package main

import "fmt"

type person struct {
	lname string
	fname string
	phone string
}

type buessness struct {
	name    string
	address string
	person
}

func (b buessness) info() {
	fmt.Printf("Contact at %s , %s %s\n", b.name, b.fname, b.person.lname)

}

func main() {
	cont1 := person{
		"Name",
		"First",
		"999-999-99-999",
	}

	cont2 := buessness{
		"StI",
		"address",
		cont1,
	}
	cont2.info()
}
