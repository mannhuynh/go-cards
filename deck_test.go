package main

import (
	"myapp/deck"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := deck.NewDeck()

	d.SavetoFile("_decktesting")

	loadedDeck := deck.NewDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
