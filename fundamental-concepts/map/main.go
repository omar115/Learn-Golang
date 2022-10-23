/*
# a map is a collection of key value pairs
- it is similar to Python dictionary
- both key and values are statically typed

# how to create an empty map
colors := make(map[string]string)
- it will create an empty map
- the value will be zero

# difference between struct and map
- map: all keys and values will be same type
- struct: all values can be different type
- in map, all keys are indexed
- in struct all keys are not indexed
- map is reference type
- struct is value type

*/

package main

import "fmt"

func main() {
	codes := map[string]string{"en": "English", "bn": "Bangla"}

	fmt.Println("dict is: ", codes, "\nlength of dict is: ", len(codes))
	fmt.Println(codes["en"])

	// delete a val
	delete(codes, "bn")
	fmt.Println("Now the new dict is: ", codes)
	// how to declare an empty map
	empty_map := make(map[string]int)
	fmt.Println(empty_map)

	// loop through the dict/map
	codes_2 := map[string]string{"en": "English", "bn": "Bangla", "hi": "Hindi"}

	for key, code := range codes_2 {
		fmt.Println(key, " => ", code)
	}

}
