package main

import "fmt"

func main() {
	var age int
	fmt.Printf("Please enter a number of your age: ")
	fmt.Scan(&age)

	switch {
	case age < 10:
		fmt.Println("You are less than 10 years")
	case age < 20:
		fmt.Println("You are more than 10 years")
	default:
		fmt.Println("Please enter something below 20")
	}
}
