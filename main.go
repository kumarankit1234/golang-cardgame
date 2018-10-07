package main

import (
	"fmt"
)

func main() {
	cards := getCard()
	fmt.Println(cards)

}

func getCard() string {
	return "Ace of spades"
}
