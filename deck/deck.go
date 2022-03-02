package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// NewDeck function create a new deck of cards
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

// PrintDeck will print out the cards of a deck
func (d deck) PrintDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Deal will split a deck into two decks
func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// DecktoString return a string by converting deck value to slice of string and joining to a string
func (d deck) DecktoString() string {
	return strings.Join([]string(d), ",")
}

// SavetoFile will save a deck to a file of string
func (d deck) SavetoFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.DecktoString()), 0666) // 0666 is a permition number

}

// NewDeckFromFile convert a file of string to a string slice with comma separation
func NewDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// Shuffle will shuffle a deck of card
func (d deck) Shuffle() {
	// Seeding the random number
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
