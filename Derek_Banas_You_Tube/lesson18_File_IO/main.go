package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//	f, err := os.Create("data.txt")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer f.Close()
	//
	//	iPrimeArr := []int{1, 2, 3, 4, 5, 8}
	//	var sPrimeArr []string
	//	for _, v := range iPrimeArr {
	//		sPrimeArr = append(sPrimeArr, strconv.Itoa(v))
	//	}
	//	for _, num := range sPrimeArr {
	//		_, err := f.WriteString(num + "\n")
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//
	//	f, err = os.Open("data.txt")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer f.Close()
	//	scan := bufio.NewScanner(f)
	//	for scan.Scan() {
	//		fmt.Println("Pritme", scan.Text())
	//	}
	//	if err := scan.Err(); err != nil {
	//		panic(err)
	//	}

	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file not found")
	} else {
		f, err := os.OpenFile("data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fmt.Println(f.Name())

		if _, err := f.WriteString("11\n"); err != nil {
			panic(err)
		}

	}

}
