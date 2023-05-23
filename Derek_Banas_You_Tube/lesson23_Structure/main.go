package main

import "fmt"

//type customer struct {
//	name    string
//	address string
//	balance float64
//}
//
//func main() {
//	var ts customer
//	ts.name = "John Doe"
//	ts.address = "lisunova 5"
//	ts.balance = 100.65
//	getCustomerIbfo(ts)
//	newCustomerAdd(&ts, "lisunova 7")
//	fmt.Println("Address:", ts.address)
//
//	sS := customer{"Salli Smith", "6Gabbq", 0.0}
//	fmt.Println("name:", sS.name)
//}
//
//func getCustomerIbfo(c customer) {
//	fmt.Printf("%s owes us %.2f\n", c.name, c.balance)
//}
//
//func newCustomerAdd(c *customer, address string) {
//	c.address = address
//}

//type rectangle struct {
//	hight, length float64
//}
//
//func (r rectangle) Arear() float64 {
//	return r.hight * r.length
//}

type address struct {
	email string
	prhne string
}

type contact struct {
	name string
	age  int
	address
}

func main() {
	//rect := rectangle{12.9, 14.5}
	//fmt.Println(rect.Arear())

	c := contact{
		name: "John",
		age:  13,
		address: address{
			email: "john@example.com",
			prhne: "998999999999",
		},
	}
	fmt.Println(c.email)
}
