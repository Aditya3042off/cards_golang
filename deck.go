package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
Create a custome type 'deck' which is a slice of strings
*/
type deck []string


func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Clubs","Spades","Diamonds","Hearts"}
	cardValues := []string{"Ace","Two","Jack","King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}


func (d deck) print() {
	
	for i,card := range d {

		fmt.Println(i,card)
	}
}


func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}


/* Convert deck to string type */
func (d deck) toString() string {

	return strings.Join([]string(d), ",")   
}


/* Save the deck to a local file */
func (d deck) saveToFile(filename string) error {

	deckString := []byte(d.toString())

	return os.WriteFile(filename, deckString, 0666)
}


/* Get a new deck from a local file */
func newDeckFromFile(filename string) deck {

	bs, err := os.ReadFile(filename)

	if err != nil {

		fmt.Println("Error reading deck from file: ", err)
		os.Exit(1)
	}

	deckString := string(bs)
	cards := deck(strings.Split(deckString, ","))

	return cards
	
}


/* Shuffle the deck */
func (d deck) shuffle() {

	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	myRand := rand.New(src) // my custom rand object

	for idx := range d {
		randIdx := myRand.Intn(len(d))
		
		d[idx], d[randIdx] = d[randIdx], d[idx]
	}

}