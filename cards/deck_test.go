package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v.", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected first card to be 'Ace of Diamonds', but got %v.", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected last card to be 'King of Spades', but got %v.", d[len(d)-1])
	}
}

func TestSaveToFileAndReadDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"

	os.Remove(testFileName)

	d := newDeck()
	d.saveToFile(testFileName)

	loadedDeck := readDeckFromFile(testFileName)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v.", len(loadedDeck))
	}

	os.Remove(testFileName)
}
