package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "ff0000",
		"white": "ffffff",
		"black": "000000",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("color => %v | hex => %v.\n", color, hex)
	}
}
