package main

import (
	"fmt"
)

func main() {
	//mnogomernieMassivi := [2][2]int{
	//	{1, 2},
	//	{3, 4},
	//}
	//
	//for i := 0; i < 2; i++ {
	//	for j := 0; j < 2; j++ {
	//		fmt.Println(mnogomernieMassivi[i][j])
	//	}
	//}
	aStr := "asdhybjhbuhybooo"
	aRa := []rune(aStr)
	for _, v := range aRa {
		fmt.Println(v)
	}

	Arr := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'g'}
	BStr := string(Arr[:])
	fmt.Println(BStr)
}
