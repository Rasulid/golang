//Операция "равно". Возвращает true, если оба операнда равны, и false, если они не равны:

package main

import "fmt"

func main() {
	var a int = 8
	var b int = 3
	var c bool = a == b
	fmt.Println(c) // false

	//>

	//Операция "больше чем". Возвращает true, если первый операнд больше второго, и false, если первый операнд меньше второго:
	//
	//a = 8
	//b = 3
	//c = a > b // true
	////<
	//
	////Операция "меньше чем". Возвращает true, если первый операнд меньше второго, и false, если первый операнд больше второго:
	//
	//a = 8
	//b = 3
	//c = a < b // false
	////<=
	//
	////Операция "меньше или равно". Возвращает true, если первый операнд меньше или равен второму, и false, если первый операнд больше второго:
	//
	//a = 8
	//b = 3
	//c = a <= b // false
	////>=
	//
	////Операция "больше или равно". Возвращает true, если первый операнд больше или равен второму, и false, если первый операнд меньше второго:
	//
	//a = 8
	//b = 3
	//c = a >= b // true
	//
	////!=
	//
	////Операция "не равно". Возвращает true, если первый операнд не равен второму, и false, если оба операнда равны:
	//
	//a = 8
	//b = 3
	//c = a != b // true
	//d := a != 8 // false
	//
	////Как правило, операции отношения применяются в условных конструкциях типа if...else, которые мы далее рассмотрим.
	////
	////Логические операции
	////Логические операции сравнивают два условия. Как правило, они применяются к отношениям и объединяют несколько операций отношения. К логическим операциям относят следующие:
	////
	////! (операция отрицания)
	////
	////Инвертирует значение. Если операнд равен true, то возвращает false, иначе возвращает true.
	//
	//a = true
	//b = !a //false
	//c = !b // true
	//
	////&& (конъюнкция, логическое умножение)
	////
	////Возвращает true, если оба операнда не равны false. Возвращает false, если хотя бы один операнд равен false.
	//
	//b = 4 > 5 && 6 > 8   //false
	//c = 3 <= 5 && 10 > 8 // true
	//
	////|| (дизъюнкция, логическое сложение)
	////
	////Возвращает true, если хотя бы один операнд не равен false. Возвращает false, если оба операнда равны false.
	//
	//b = 4 > 5 || 6 > 8   //false
	//c = 3 == 5 || 10 > 8 // true

}
