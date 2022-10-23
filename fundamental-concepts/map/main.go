package main

import "fmt"

func main() {
	codes := map[string]string{"en": "English", "bn": "Bangla"}

	fmt.Println(codes)

	// how to declare an empty map
	empty_map := make(map[string]int)
	fmt.Println(empty_map)

}
