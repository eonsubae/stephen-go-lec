package main

func main() {
	cards := NewDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
