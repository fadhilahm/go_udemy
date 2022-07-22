package main

import "fmt"

func main() {
	cards := newDeck()

	// hands, remainingCards := deal(cards, 5)

	// hands.print()
	// remainingCards.print()

	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")
}
