package main

func main() {

	cards := newDeck()

	//cards.saveToFile("harish_cards")

	//cards := newDeckFromFile("harish_card")
	cards.shuffle()
	cards.print()

}
