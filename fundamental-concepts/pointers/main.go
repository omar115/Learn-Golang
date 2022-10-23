/*
a pointer is a variable that holds memory address of
another variable.
They point where the memory is allocated and
provide ways to find or even change the value located at the memory location.

*/

package main

import "fmt"

func main() {

	s := "hello"

	var b *string = &s
	fmt.Println(b)

}
