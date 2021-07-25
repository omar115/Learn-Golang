package main

import "fmt"

func main() {
	var_list()
}

func var_list() {

	var ages = [3]int{20, 25, 30}
	fmt.Println("The array of ages is ", ages)
	fmt.Println("The length of array is ", len(ages))

	names := [4]string{"yoshi", "mern", "pich", "jason"}
	fmt.Println("The array of names is ", names)

	// append values to an array
	scores := []int{100, 200, 300}
	new_list := append(scores, 400)
	fmt.Println("After appending data, the new list is: ", new_list)

	// slice array
	slice_var := new_list[1:3]
	slice_var2 := new_list[1:]
	slice_var3 := new_list[:3]
	fmt.Println(slice_var)
	fmt.Println(slice_var2)
	fmt.Println(slice_var3)

}
