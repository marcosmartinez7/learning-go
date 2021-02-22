package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// A deck is represented by a slice of strings
// Here we create a new type that "extends" the slice of strings
type deck []string

//////////////////////
// Deck  functions //
/////////////////////

// Constructor of a deck
func newDeck() deck {
	cardSuits, cardValues := getSuitsAndValues()
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
	minDealingHandPosition := getRandomNumber(len(deckValue))
	maxDealingHandPosition := minDealingHandPosition + handSize
	dealHand := deckValue[minDealingHandPosition:maxDealingHandPosition]
	firstDeckHalf := deckValue[:minDealingHandPosition]
	secondDeckHalf := deckValue[maxDealingHandPosition:]
	*d = append(firstDeckHalf, secondDeckHalf...)
	return dealHand
}

// Saves the current deck into a file.
func (d deck) saveToFile(fileName string) error {
	deckBytes := d.toBytesSlice()
	return ioutil.WriteFile(fileName, deckBytes, 0666)
}

// Restrores the current deck from a file
func newDeckFromFile(filename string) (deck, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return byteSliceToDeck(content), nil
}

// Shuffles the deck. This functions implements the Fisher–Yates shuffle algorithm
func (d *deck) shuffle() {
	for i := len(*d) - 1; i > 0; i-- {
		fmt.Println("thge i", i)
		j := getRandomNumber(i)
		aux := (*d)[i]
		(*d)[i] = (*d)[j]
		(*d)[j] = aux
	}
}

/////////////////////////
// Auxiliary functions //
/////////////////////////

// String representation of a deck
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Converts a deck into a byte slice
func (d deck) toBytesSlice() []byte {
	return []byte(d.toString())
}

// Converts a byte slice into a deck
func byteSliceToDeck(deckBytes []byte) deck {
	stringDeck := string(deckBytes)
	return deck(strings.Split(stringDeck, ","))
}

// Gets a random int between 0 and maxValue
func getRandomNumber(maxValue int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxValue)
}

// Constant-like function that returns the suits and values of a deck
func getSuitsAndValues() ([]string, []string) {
	cardSuits := []string{"Diamonds", "Hearths", "Spades", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}
	return cardSuits, cardValues
}
