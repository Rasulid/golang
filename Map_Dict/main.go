package main

func main() {
	var studets = map[int]string{
		1: "Rasul",
		2: "Jovo",
		3: "Sul",
		4: "Vendas",
		5: "Nove",
	}
	studets[5] = "Ogooooo"
	delete(studets, 4)
	for k, v := range studets {
		println(k, v)
	}
}
