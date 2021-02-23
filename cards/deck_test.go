package main

import (
	"os"
	"testing"
)

func TestNewDeckSize(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(deck))
	}
}

func TestNewDeckFirstCard(t *testing.T) {
	deck := newDeck()
	if deck[0] != "Ace of Diamonds" {
		t.Errorf("Expected first card Ace of Diamonds, but got %v", deck[0])

	}
}

func TestNewDeckLastCard(t *testing.T) {
	deck := newDeck()
	if deck[len(deck)-1] != "Two of Clubs" {
		t.Errorf("Expected last card Two of Clubs, but got %v", deck[len(deck)-1])

	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("decktesting")
	deck := newDeck()
	deck.saveToFile("decktesting")
	loadedDeck, _ := newDeckFromFile("decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(loadedDeck))
	}
	os.Remove("decktesting")
}
