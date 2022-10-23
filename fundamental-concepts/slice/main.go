/*

slice is similar concept to python list.
Similar to list, slice can be extended to any length,

slice can do:
- add items
- delete items
- copy items
- make new slice

*/

package main

import "fmt"

func main() {
	// how to declare a slice?
	nums := []int{10, 20, 30}

	// how to add new entry into the slice ?
	nums = append(nums, 40, 50)

	// how to run a for loop of a slice?
	for i, num := range nums {
		fmt.Printf("slice index: %d || slice index value: %d \n", i, num)
	}

	// subslice
	nums_2 := nums[0:3]

	fmt.Println(nums_2)
	fmt.Println(len(nums_2))
	fmt.Println(cap(nums_2))

	// how to copy the slice
	src_slice := []int{1, 2, 33}
	// dst_slice := make([]int, 3) // will make a new slice with default [0,0,0]
	dst_slice_2 := []int{4, 5, 6}
	rst := copy(dst_slice_2, src_slice)
	fmt.Println("nums of elements copied: ", rst)
	fmt.Println("destination slice is: ", dst_slice_2)

}
