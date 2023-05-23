package main

import "fmt"

// func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращаемых_результатов){
//    тело_метода
//}

type person struct {
	name string
	age  int
}

type library []string

func (l library) print() {
	for _, v := range l {
		fmt.Println(v)
	}
}

func main() {
	var lib library = library{"books1", "books2", "books3", "books4"}
	lib.print()
	var tom = person{name: "Tom", age: 28}
	tom.man()
	tom.eat("plow")

}

func (p person) man() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
}

func (p person) eat(d string) {
	fmt.Println(p.name, "eat", d)
}
