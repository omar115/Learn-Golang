package main

import "fmt"

func main() {

	var a [3]int
	fmt.Println("array: ", a)

	b := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println("array 2nd type: ", b)
	fmt.Println("length: ", len(b))

	// run loop over the array
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// run loop over the array (2nd way)
	for index, element := range b {
		fmt.Println(index, " => ", element)
	}

	// multidimensional array
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr[2][1])

}
