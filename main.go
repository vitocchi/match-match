package main

import (
	"fmt"

	"github.com/vitocchi/neurathenia/card"
)

func main() {
	card, err := card.NewCard(card.Spade, 1)
	fmt.Printf("suit = %+v\n", card)
	fmt.Printf("error = %+v\n", err)
}
