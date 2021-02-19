package main

import "fmt"

// A deck is represented by a slice of strings
// Here we create a new type that "extends" the slice of strings

type deck []string

// Function with deck as a receiver
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
