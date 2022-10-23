package main

import "fmt"

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum

}

func main() {
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(1, 2, 3))

}
