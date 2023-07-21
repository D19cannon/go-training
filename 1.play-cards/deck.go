package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// create a new deck of type deck
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

// receiver function type.print()
// can be used like deck.print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
