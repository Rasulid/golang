package main

import "fmt"

//func main() {
//	var pArr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	dbArrey(&pArr)
//	fmt.Println(pArr)
//
//}
//
//func dbArrey(arr *[10]int) {
//	for i := 0; i < len(arr); i++ {
//		arr[i] *= 2
//	}
//}

func main() {
	iSlice := []float64{11, 13, 17}
	fmt.Printf("%.3f\n", getAvarege(iSlice...))
}

func getAvarege(nums ...float64) float64 {
	var sum float64 = 0.0
	var sumSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return sum / sumSize
}
