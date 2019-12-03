package main

import (
	"math/rand"
	"time"

	"github.com/vitocchi/match-match/table"
	"github.com/vitocchi/match-match/table/card"
)

type RandomStrategy struct{}

func (s *RandomStrategy) PickCards(cs card.Cards) [2]card.Card {
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
	p := make([]table.Player, 1, 1)
	p[0] = table.NewPlayer(&RandomStrategy{})
	rand.Seed(time.Now().UnixNano())
	table := table.NewTable(p)
	table.ExecGame()
}
