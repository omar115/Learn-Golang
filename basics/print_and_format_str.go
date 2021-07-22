package main

import "fmt"

func main() {
	testing()
}

func testing() {
	age := 25
	name := "Omar"

	fmt.Println("Hi! My name is ", name, " and I am ", age, " years old.")

	// formated string using format specifier
	// %_ = format specifier
	// %v means variables, %q means quotes, %T means types, %f means float
	// to print in one/two/three, etc. decimal points use %0.1f
	fmt.Printf("Hello! My name is %v and I am %v years old. Nice to meet you! \n", age, name)
	fmt.Printf("The type of age is %T \n", age)
	fmt.Printf("My Score was %0.01f points! \n", 23232.2312)

}
