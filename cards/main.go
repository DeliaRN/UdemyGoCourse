package main

// Code to create and manipulate the deck itself

func main() {

	//making a slice of cards:
	cards := newDeck()
	/*
		TESTING FUNCTIONS:

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
	*/

	cards.saveToFile("my_deck")

}

func newCard() string {
	return "Five of Diamonds"
}
