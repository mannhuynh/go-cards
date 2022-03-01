package main

import (
	"fmt"
	"myapp/deck"
)

func main() {
	cards := deck.NewDeck()

	hand, remainDeck := deck.Deal(cards, 5)
	hand.PrintDeck()
	fmt.Println("-----")
	remainDeck.PrintDeck()

}
