package main

import "fmt"

func main() {
	cards := deck{"Hello", "Hello!"}
	fmt.Println(cards)

	cards.print()
}
