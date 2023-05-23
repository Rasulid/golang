package main

import "fmt"

func main() {
	/*
			for [инициализация счетчика]; [условие]; [изменение счетчика]{
		    // действия
		}
	*/

	for i := 1; i < 10; i++ {
		//fmt.Printf("%Method:", i)
		//fmt.Println(i * i)
		//fmt.Println("================================")
	}

	i := 1
	for i < 10 {
		//fmt.Println(i)
		i++
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
			if i == 4 || j == 4 {
				break
			}
		}
		fmt.Println()

	}

	/*
		Для перебора массивов можно использовать следующую форму цикла for:
		for индекс, значение := range массив{
		    // действия
		}
	*/

	for _, value := range []string{"TOm", "Serik", "Berik"} {
		if value == "Serik" {
			break
		}
		fmt.Println(value)
	}
}
