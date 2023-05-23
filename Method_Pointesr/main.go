package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) updateAge(newAge int) {
	p.age = newAge
}
func main() {
	var TomPerson = person{name: "Tom", age: 12}
	fmt.Println("before: ", TomPerson.age)
	TomPerson.updateAge(151)
	fmt.Println("after: ", TomPerson.age)

	var TomPointer *person = &TomPerson
	fmt.Println("before: ", TomPointer.age)
	TomPerson.NewupdateAge(151)
	fmt.Println("after: ", TomPointer.age)

}

func (p *person) NewupdateAge(newAge int) {
	(*p).age = newAge
}
