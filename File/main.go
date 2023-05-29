package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := "Text to file"
	file.WriteString(data)

	file_data, err := os.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file_data))
	defer os.Remove(file.Name())

}
