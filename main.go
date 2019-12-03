package main

import (
	"math/rand"
	"time"

	"github.com/vitocchi/match-match/table/card"
)

type RandomStrategy struct{}

func (s *RandomStrategy) pickCards(cs card.Cards) [2]card.Card {
	first := rand.Intn(len(cs))
	var second int
	for {
		second = rand.Intn(len(cs))
		if first != second {
			break
		}
	}
	return [2]card.Card{cs[first], cs[second]}
}

func main() {
	p := make([]card.Player, 1, 1)
	p[0] = card.NewPlayer(&RandomStrategy{})
	rand.Seed(time.Now().UnixNano())
	table := card.NewTable()
	table.ExecGame()
}
