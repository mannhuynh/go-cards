package deck

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}
	return cards
}

func (d deck) PrintDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// DecktoString return a string by converting deck value to slice of string and joining to a string
func (d deck) DecktoString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) SavetoFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.DecktoString()), 0666)

}
