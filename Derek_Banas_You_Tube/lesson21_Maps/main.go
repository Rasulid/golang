package main

import "fmt"

func main() {
	// var name = make(map[keyType]valueType)
	var heroes map[string]string
	heroes = make(map[string]string)
	villians := make(map[string]string)
	heroes["Betman"] = "Bruse Wain"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"

	villians["Lex Luter"] = "Lex Luter"

	supperPets := map[int]string{1: "cat", 2: "dog", 3: "horse"}

	fmt.Printf("%v is Batman\n", heroes["Betman"])

	fmt.Println("cat is ", supperPets[1])

	delete(heroes, "The Flash")

	for k, v := range heroes {
		fmt.Printf("%v:  %v\n", k, v)
	}
}
