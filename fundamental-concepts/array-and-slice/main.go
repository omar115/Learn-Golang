package main

import "fmt"

func main() {
	// how to declare a slice?
	cards := []string{newCard(), "Omar!"}

	// how to add new entry into the slice ?
	cards = append(cards, "How are you ?")

	// how to print the slice ?
	// fmt.Println(cards)

	// how to run a for loop of a slice?
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "AsSalamu Alaikum, "
}
