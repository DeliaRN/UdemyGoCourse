package main

// Code to create and manipulate the deck itself

func main() {

	//making a slice of cards:

	/*
		TESTING FUNCTIONS:

		CREATE A NEW DECK:
			cards := newDeck()
			cards.print()

		DEAL FUNCTION:
			hand, remainingCards := deal(cards, 5)

			hand.print()
			fmt.Println("---")
			remainingCards.print()

			greeting := "Hi there!"
			fmt.Println([]byte(greeting))

		TO STRING FUNCTION:
		stringedCards := cards.toString()
		fmt.Println(stringedCards)

		SAVE TO FILE FUNCTION:
		cards.saveToFile("my_deck")

		OPEN EXISTING DECK:
		cards := newDeckFromFile("my_deck")
		cards.print()
	*/

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
