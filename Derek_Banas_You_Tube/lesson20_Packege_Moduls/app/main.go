package main

import (
	stuff "exaple/project/mypackege"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello", stuff.Name)
	ss := []int{1, 2, 3, 4}
	sF := stuff.IntArrToStrArr(ss)
	fmt.Println(sF)
	fmt.Println(reflect.TypeOf(sF))

}
