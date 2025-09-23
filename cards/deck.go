package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Any variable of type deck now has access to all the methods we define on it
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// convert deck to []string itself, even though deck is a slice of strings
	// use Join from strings package to join the elements of the slice into a single string, with commas in between
}
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
	// saves in a plain text file within the same folder as our code
	// 0666 gives read and write permissions to the file for everyone
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) //byteslice, error
	if err != nil {
		// option #1: log the error and return a call to newDeck()
		// option #2: log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	// deck(strings.Split(string(bs), ","))
	return deck(s)
}

func (d deck) shuffle() {

	for i := range d {

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		// newPosition := rand.Intn(len(d) - 1)
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
		// Fancy swap line
	}

}
