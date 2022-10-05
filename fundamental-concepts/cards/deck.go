package main

import "fmt"

// create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Carroms"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards

}

// d is the reference to deck, you can
// write anything you want

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// here d is representating the deck type
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}
