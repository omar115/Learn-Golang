package main

import "fmt"

func addNumbers(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("addNumbers: ", addNumbers(2, 3))
}
