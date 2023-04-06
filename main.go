package main

func main() {
	cards := newDeck()
	cards.saveToFile("deck")

	cards.shuffle()

	cardsFromFile := newDeckFromFile("deck")
	cardsFromFile.print()

}
