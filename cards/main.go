package main

import (
	"flag"
	"log"
)

func main() {

	load := flag.Bool("load", false, "load game")
	flag.Parse()
	log.Println(*load)

	currentDeckLocation := "currentDeck"
	currentHandLocation := "currentHand"

	// Initialize deck
	var cards deck
	if *load {
		cards, _ = newDeckFromFile("currentDeck")

	} else {
		cards = newDeck()
	}

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
	cards.saveToFile(currentDeckLocation)

	// Save current hand to file
	log.Println("[INFO] Saving current hand to file: " + currentHandLocation)
	currentHand.saveToFile(currentHandLocation)

}
