package main

import (
	"myapp/deck"
)

func main() {

	cards := deck.NewDeck()

	// hand, remainDeck := deck.Deal(cards, 5)
	// fmt.Println(hand.DecktoString()
	// remainDeck.PrintDeck()

	cards.SavetoFile("mycards")

	cardsFromFile := deck.NewDeckFromFile("mycards")
	// cardsFromFile.PrintDeck()

	cardsFromFile.Shuffle()
	cardsFromFile.PrintDeck()

}
