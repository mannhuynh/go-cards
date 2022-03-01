package main

import (
	"myapp/deck"
)

func main() {
	cards := deck.NewDeck()

	cards.SavetoFile("mycards.txt")

	// hand, remainDeck := deck.Deal(cards, 5)
	// fmt.Println(hand.DecktoString())
	// fmt.Println("-----")
	// remainDeck.PrintDeck()

}
