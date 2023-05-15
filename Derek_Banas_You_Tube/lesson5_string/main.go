package main

import (
	"fmt"
	"strings"
)

func main() {
	// Replace

	sV1 := "A Word"
	replcer := strings.NewReplacer("A", "Another")
	sV2 := replcer.Replace(sV1)
	fmt.Println(sV2)

	// Len
	fmt.Println("Len :", len(sV2))

	//Contains
	fmt.Println("Contains :", strings.Contains(sV2, "Another"))

	//Index
	fmt.Println("Index :", strings.Index(sV2, "o"))

	//Raplace
	fmt.Println("Raplace :", strings.Replace(sV2, "o", "0", -1))

	//Split
	fmt.Println("Split :", strings.Split(sV2, "o"))

	//ToLower
	fmt.Println("ToLower :", strings.ToLower(sV2))

	//ToUpper
	fmt.Println("ToUpper :", strings.ToUpper(sV2))
}
