package main

import "fmt"

type Animal interface {
	AngrySound()
	HappySound()
}
type Cat string

func (c Cat) Attack() {
	fmt.Println("Cat attack")
}
func (c Cat) Name() string {
	return string(c)
}
func (c Cat) AngrySound() {
	fmt.Println("Cat says hisssss")
}
func (c Cat) HappySound() {
	fmt.Println("Cat says murrrrr")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	fmt.Println("Kitty name:", kitty2.Name())
}
