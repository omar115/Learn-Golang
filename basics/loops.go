package main

import "fmt"

func main() {
	var_list()
}

func var_list() {

	// for x := 0; x < 5; x++ {

	// 	fmt.Println("The value of x is : ", x)

	// }

	names := []string{"omar", "sajid", "mojnu", "hasib", "jagi", "bepari"}

	for i := 0; i < len(names); i++ {
		fmt.Println("The value of index ", i, " is : ", names[i])
	}

	// alternative approach

}
