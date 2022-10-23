/*

a defer statement delays the execution of a function until
the surrounding function retruns

to sum up: if you use defer statement before a function, the function
will be executed last after all the function

*/

package main

import "fmt"

func printName(name string) {
	fmt.Println("name: ", name)
}

func rollNumber(n int) {
	fmt.Println("rollNumber: ", n)
}

func address(a string) {
	fmt.Println("address: ", a)
}

func main() {
	defer printName("biden")
	rollNumber(12)
	address("Paikpara, Manikdi")
}
