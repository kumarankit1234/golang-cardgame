package main

func main() {

	cards := deck{"Ace of diamond", getCard()}
	cards.print()
}

func getCard() string {
	return "Ace of spades"
}
