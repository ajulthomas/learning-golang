package main

import "fmt"

// deck is a slice of strings
type deck []string

func newDeck() deck {

	newDeck := deck{}

	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			card := value + " of " + suite
			newDeck = append(newDeck, card)
		}
	}

	return newDeck

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
