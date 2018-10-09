package main

import (
	"fmt"
)

func main() {

	cards := newDeckFromFile("/Users/ankit.sin/my-cards.txt")
	fmt.Println(cards.toString())

}
