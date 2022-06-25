package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	first_name string
	last_name  string
	contact    contactInfo
}

func main() {
	jim := person{
		first_name: "Jim",
		last_name:  "Donalt",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 1205,
		},
	}
	// fmt.Println(jim)
	fmt.Printf("%+v", jim)
}
