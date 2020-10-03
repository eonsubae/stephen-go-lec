package main

func main() {
	cards := NewDeck()
	cards.saveToFile("my_cards")
}

func newCard() string {
	return "Five of Diamonds"
}
