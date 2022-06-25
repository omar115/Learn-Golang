package main

import (
	"fmt"
)

func main() {
	// declare a map
	colors := map[string]string{
		"red":  "#ff0000",
		"blue": "#ff0001",
	}
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is >>", hex)
	}
}
