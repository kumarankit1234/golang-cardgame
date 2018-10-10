package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(cards))
	}

	if cards[0] != "Ace of Diamond" {
		t.Errorf("Expected first card to be Ace of Diamond but got %v", cards[0])
	}

	if cards[len(cards)-1] != "King of Clubs" {
		t.Errorf("Expected last cards to be King of Clubs but got %v", cards[len(cards)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	cards := newDeck()
	cards.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
