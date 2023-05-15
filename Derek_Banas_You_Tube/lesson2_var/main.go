package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		var name type
		Если переменная написана в нижнем регистре то её статус считается private
					Это заначит его нельзя импортировать
		А Если имя переменного начинается Вверхнеем регистре то она считается
					public
	*/

	/*
		var vName string = "Rasul"
		var v1, v2 = 2.4, 4.3
		var v3 = "Hello, world"
		v4 := 123.4
		v4 = 23
		//v4 := 29  !!!Error
	*/

	fmt.Println(reflect.TypeOf(12))
	fmt.Println(reflect.TypeOf("Hello, world"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(123.4))

}
