package main

import "fmt"

func main() {
	// Initialize deck
	cards := newDeck()
	fmt.Println("[INFO] Full deck size: ", len(cards))

	// Hand dealing
	fmt.Println("[INFO] Dealing cards: ")
	dealingCards := deal(&cards, 2)
	dealingCards.print()
	fmt.Println("")
	fmt.Println("[INFO] Current deck ")
	cards.print()
	fmt.Println("[INFO]  Current deck size", len(cards))
}
