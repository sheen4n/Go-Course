package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
