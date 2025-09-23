package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// delete the test deck if it already exists

	os.Remove("_decktesting")

	// create a deck
	deck := newDeck()

	// save to file _decktesting
	deck.saveToFile("_decktesting")

	// load from file
	loadedDeck := newDeckFromFile("_decktesting")

	// assert deck length
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 53, but got %v", len(loadedDeck))
	}
	// delete any files in current working directory with the name _decktesting
	os.Remove("_decktesting")

}
