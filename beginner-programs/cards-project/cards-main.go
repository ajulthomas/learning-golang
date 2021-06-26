package main

func main() {
	deck := newDeck()
	hand, remainingCards := deal(deck, 4)
	hand.print()
	remainingCards.print()
}
