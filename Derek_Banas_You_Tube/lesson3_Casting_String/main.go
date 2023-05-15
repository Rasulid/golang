package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	Cv1 := "500000" // str to int
	Cv2, err := strconv.Atoi(Cv1)
	println(Cv2, err, reflect.TypeOf(Cv2))

	Cv3 := "12.4" // str to float64 - float32
	if Cv4, err := strconv.ParseFloat(Cv3, 64); err == nil {
		fmt.Println(Cv4)
	}

	Cv5 := fmt.Sprintf("%f", 1234.5)
	println(Cv5, reflect.TypeOf(Cv5))
}
