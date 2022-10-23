/*
# Description of Struct

* struct is a data structure
* it is a collection of properties that are
    - related together
* it is a similar concept of Dictionary of Python

user-defined data type.
• a structure that groups together data elements.
• provide a way to reference a series of grouped values through a single variable name.
• used when it makes sense to group or associate two or more data variables.

*/

package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius int
}

func main() {
	var c Circle
	c.x = 5
	c.y = 5
	c.radius = 5
	fmt.Printf("%+v \n", c)
}
