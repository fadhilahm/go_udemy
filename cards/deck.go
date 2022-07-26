package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func readDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(byteSlice), ",")
}

func (d deck) shuffle() {
	maxNumber := len(d)
	for i := range d {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomIndex := r.Intn(maxNumber)
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}
