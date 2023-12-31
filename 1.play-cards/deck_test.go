package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expetected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card to be 'Four of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	prefix := "_decktesting"
	os.Remove(prefix)

	deck := newDeck()
	deck.saveToFile(prefix)

	loadDeck := newDeckFromFile(prefix)

	if len(loadDeck) != 16 {
		t.Errorf("Expecting 16 cards in deck, got %v", len(loadDeck))
	}

	os.Remove(prefix)
}
