package main

import (
	"log"
	"strconv"
)

func main() {
	currentDeckLocation := "currentDeck"
	currentHandLocation := "currentHand"

	// Initialize deck
	cards := newDeck()
	log.Println("[INFO] Full deck size: ", len(cards))

	// Hand dealing
	log.Println("[INFO] Dealing cards: ")
	currentHand := cards.deal(2)
	currentHand.print()
	log.Println("[INFO] Current deck ")
	cards.print()
	log.Println("[INFO] Current deck size", len(cards))

	// Save current game deck to file
	log.Println("[INFO] Saving current deck to file: " + currentDeckLocation)
	currentDeckSaved := cards.saveToFile(currentDeckLocation)
	log.Println("[INFO] Current deck saved: " + strconv.FormatBool(currentDeckSaved))

	// Save current hand to file
	log.Println("[INFO] Saving current hand to file: " + currentHandLocation)
	currentHandSaved := currentHand.saveToFile(currentHandLocation)
	log.Println("[INFO] Current hand saved: " + strconv.FormatBool(currentHandSaved))

}
