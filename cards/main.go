package main
// to run go run main.go deck.go
func main() {

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

