package main

import "fmt"

func main() {
	cards := NewDeck()
	fmt.Println(cards.toString())
}

func newCard() string {
	return "Five of Diamonds"
}
