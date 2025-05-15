package main


func main() {

	// cards := newDeck()

	// hand, remainingCards := deal(cards,5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// err := cards.saveToFile("my_cards")

	// if err != nil {

	// 	fmt.Println("Error saving deck to file: ", err)
	// }

	// cards := newDeckFromFile("my_cards")

	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}