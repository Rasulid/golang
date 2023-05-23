package main

import "fmt"

type Tsp float64
type Tbs float64
type ML float64

func tspToMl(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func tbsToMl(tsp Tbs) ML {
	return ML(tsp * 14.79)
}

func main() {
	ml1 := ML(tspToMl(3) * 4.92)
	fmt.Printf("ml1 = %.2f\n", ml1)

	ml2 := ML(tbsToMl(3) * 14.79)
	fmt.Printf("ml1 = %.2f\n", ml2)
}
