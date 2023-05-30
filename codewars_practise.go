package main

import "fmt"

func main() {
	fmt.Println(HowManyDalmatians(80))
	fmt.Println(GetPlanetName(90))
	fmt.Println(Namevar())
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
