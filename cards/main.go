package main

import "fmt"

func main() {
	cards := deck{"Hello", "Hello!"}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	cards.print()
}
