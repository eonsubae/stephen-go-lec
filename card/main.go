package main

import "fmt"

func main() {
	// cards := NewDeck()

	// hand, remainingDeck := deal(cards, 5)

	// fmt.Println("Print start")
	// hand.print()
	// fmt.Println("Hand is printed")
	// remainingDeck.print()
	// fmt.Println("Remaining is printed")

	greeting := "Hi there!"
	fmt.Println([]byte(greeting))
}

func newCard() string {
	return "Five of Diamonds"
}
