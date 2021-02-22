package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// A deck is represented by a slice of strings
// Here we create a new type that "extends" the slice of strings
type deck []string

// Constructor of a deck
func newDeck() deck {
	cardSuits := []string{"Diamonds", "Hearths", "Spades", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}
	deck := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			var cardBuilder strings.Builder
			cardBuilder.WriteString(value)
			cardBuilder.WriteString(" of ")
			cardBuilder.WriteString(suit)
			deck = append(deck, cardBuilder.String())
		}
	}
	return deck
}

// Prints the full deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

//Returns a hand from the deck and remove them from the original deck
func deal(dp *deck, handSize int) deck {
	d := *dp
	deckSize := len(d)
	rand.Seed(time.Now().UnixNano())
	minDealingHandPosition := 1 + rand.Intn(deckSize-1)
	maxDealingHandPosition := minDealingHandPosition + handSize
	dealDeck := d[minDealingHandPosition:maxDealingHandPosition]
	firstDeckHalf := d[:minDealingHandPosition]
	secondDeckHalf := d[maxDealingHandPosition:]
	*dp = append(firstDeckHalf, secondDeckHalf...)
	return dealDeck
}
