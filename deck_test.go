package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected another deck length, yours %v ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected Five of Clubs, but got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckFromfile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
