/*

- An interface specifies a method set and is a powerful way to introduce modularity in Go.
- Interface is like a blueprint for a method set.
- They describe all the methods of a method set by providing the function signature for each method.
- They specify a set of methods, but do not implement them.

## implementing an interface

A type implements an interface by implementing its methods.
The go language interfaces are implemented implicitly.
And it does not have any specific keyword to implement an interface.

*/

package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

type rect struct {
	length, breadth float64
}

func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {

	r := rect{length: 3, breadth: 4}
	c := square{side: 5}
	printData(r)
	printData(c)

}
