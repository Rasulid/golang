package main

import "fmt"

func main() {
	Sl1 := make([]string, 6)
	Sl1[0] = "Ra"
	Sl1[1] = "La"
	Sl1[2] = "Ma"
	Sl1[3] = "Sa"
	Sl1[4] = "So"
	for _, v := range Sl1 {
		println(v)
	}

	Sl2 := []int{
		1,
		2,
		3,
		4,
		5,
		6,
	}
	Sl3 := Sl2[0:2]
	fmt.Println("Sl2[0:2]", Sl2[:3])
	fmt.Println("Sl3[last]", Sl2[2:])
	Sl2[0] = 10
	fmt.Println("sl3", Sl3)
	Sl3[0] = 1
	fmt.Println(Sl2)

	Sl3 = append(Sl3, 12312)
	fmt.Println(Sl3)
	fmt.Println(Sl2)
}
