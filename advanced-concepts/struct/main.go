package main

import "fmt"

type person struct {
	first_name string
	last_name  string
}

func main() {
	alex := person{first_name: "Alex", last_name: "Anderson"}
	fmt.Println("After Initial assign----> ")
	fmt.Println(alex)
	alex.first_name = "Md Omar"
	alex.last_name = "Hasan"
	fmt.Println("After updating values----> ")
	fmt.Println(alex)
}
