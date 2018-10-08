package main

import "fmt"

type deck []string

func newDeck() deck {
	cardSuits := []string{"Diamond", "Spades", "Hearts", "Clubs"}
	cardNumbers := []string{"Ace", "Douce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}

	for _, cardSuit := range cardSuits {
		for _, cardNumber := range cardNumbers {
			cards = append(cards, cardNumber+" of "+cardSuit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
