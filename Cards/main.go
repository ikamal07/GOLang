package main

func main() {
	//cards := newDeck()

	// cards.print()
	// hands, remainingCards := deal(cards, 5)
	// hands.print()
	// remainingCards.print()
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	//cards.print()
	cards.shuffle()
	cards.print()
}
