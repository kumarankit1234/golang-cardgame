package main

func main() {

	cards := newDeckFromFile("/Users/ankit.sin/my-cards.txt")
	cards.shuffle()
	cards.print()

}
