package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// A deck is represented by a slice of strings
// Here we create a new type that "extends" the slice of strings
type deck []string

/////////////////////////////
// Deck receiver functions //
////////////////////////////

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
		log.Println(card)
	}
}

//Returns a hand from the deck and remove them from the original deck
func (d *deck) deal(handSize int) deck {
	deckValue := *d
	deckSize := len(deckValue)
	rand.Seed(time.Now().UnixNano())
	minDealingHandPosition := 1 + rand.Intn(deckSize-1)
	maxDealingHandPosition := minDealingHandPosition + handSize
	dealDeck := deckValue[minDealingHandPosition:maxDealingHandPosition]
	firstDeckHalf := deckValue[:minDealingHandPosition]
	secondDeckHalf := deckValue[maxDealingHandPosition:]
	*d = append(firstDeckHalf, secondDeckHalf...)
	return dealDeck
}

/////////////////////////
// Auxiliary functions //
/////////////////////////

// String representation of a deck
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Saves the current deck into a file.
func (d deck) saveToFile(fileName string) bool {
	deckBytes := d.toBytesSlice()
	err := ioutil.WriteFile(fileName, deckBytes, 0666)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Converts a deck into a byte slice
func (d deck) toBytesSlice() []byte {
	return []byte(d.toString())
}
