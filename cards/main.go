package main

import (
	"log"
	"strconv"
)

func main() {
	currentGameLocation := "currentGame"
	// Initialize deck
	cards := newDeck()
	log.Println("[INFO] Full deck size: ", len(cards))

	// Hand dealing
	log.Println("[INFO] Dealing cards: ")
	dealingCards := cards.deal(2)
	dealingCards.print()
	log.Println("[INFO] Current deck ")
	cards.print()
	log.Println("[INFO] Current deck size", len(cards))

	// Save current game to file
	log.Println("[INFO] Saving to file: " + currentGameLocation)
	gameSaved := cards.saveToFile(currentGameLocation)
	log.Println("[INFO] Game saved: " + strconv.FormatBool(gameSaved))

}
