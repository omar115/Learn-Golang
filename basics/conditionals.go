package main

import "fmt"

func main() {
	var_list()
}

func var_list() {

	num := 50

	if num < 30 {
		fmt.Println("The number is less than 30")
	} else {
		fmt.Println("The number is less than 100")
	}

}
