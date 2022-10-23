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
