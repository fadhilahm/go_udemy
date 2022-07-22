package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	cardSuits := []string{
		"Diamonds",
		"Clubs",
		"Hearts",
		"Spades",
	}

	deck := []string{}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			deck = append(deck, cardValue+" of "+cardSuit)
		}
	}

	return deck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
