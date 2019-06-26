package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length, but got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected first card will be Ace of Diamonds, but got %v", d[0])
	}
	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected first card will be Five of Clubs, but got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 20 {
		t.Errorf("Expected first card will be Ace of Diamonds, but got %v", d[0])
	}

	os.Remove("_decktesting")

}
