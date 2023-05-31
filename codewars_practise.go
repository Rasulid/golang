package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(HowManyDalmatians(80))
	fmt.Println(GetPlanetName(90))
	fmt.Println(Namevar())
	var arr = []string{"1:0", "2:0", "3:0", "4:4", "2:2", "3:3", "1:4", "2:3", "2:4", "3:4"}
	fmt.Println(Points(arr))
	fmt.Println(Litres(0.82))
	fmt.Println(Hero(10, 5))
}

func HowManyDalmatians(number int) string {
	dogs := []string{"Hardly any", "More than a handful!", "Woah that's a lot of dogs!", "101 DALMATIONS!!!"}

	if number <= 10 {
		return dogs[0]
	} else if number <= 50 {
		return dogs[1]
	} else if number >= 101 {
		return dogs[3]
	} else {
		return dogs[2]
	}

}

func GetPlanetName(ID int) string {
	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"

	}
	return "Pluto"
}

func Namevar() string {
	var a string = "code"
	var b string = "wa.rs"
	var name string = a + b
	return name
}

func Points(games []string) int {

	result := 0

	for _, game := range games {
		str := strings.Split(game, ":")
		x, y := str[0], str[1]

		switch {
		case x > y:
			result += 3
		case x == y:
			result += 1
		}
	}
	return result
}

func Litres(time float64) int {
	//The fun begins here.
	result := time * 0.5

	return int(result)
}

func Hero(bullets, dragons int) bool {
	// your code
	multDrago := dragons * 2
	fmt.Println(multDrago)

	if bullets >= multDrago {
		return true
	} else {
		return false
	}

}
