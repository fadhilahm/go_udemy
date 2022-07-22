package main

func main() {
	cards := readDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
