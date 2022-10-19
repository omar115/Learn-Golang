/*

%v -> default value
%t = type of the value
d = integers
c = characters
q = quoted characters
s = plain string
T = true or false
f = float
.2f = float upto 2 decimal places

*/

package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "arham"
	age := 6
	is_boy := true
	weight := 30.915

	fmt.Print("Name: ", name, "\nAge: ", age, "\nWeight: ", weight, "\nIsboy: ", is_boy)

	// type check of variables
	fmt.Printf("type of Name: %v \n", reflect.TypeOf(name))
	fmt.Printf("\nType of age: %v \n", reflect.TypeOf(age))
	fmt.Printf("\nType of weight: %v \n", reflect.TypeOf(weight))
	fmt.Printf("\nType of isboy: %v \n", reflect.TypeOf(is_boy))

}
