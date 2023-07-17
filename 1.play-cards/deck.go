package main

import "fmt"

// create a new deck of type of decl
// which is a slice of type string[]

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", " Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
