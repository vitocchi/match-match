package strategy

import (
	"math/rand"

	"github.com/vitocchi/match-match/table/card"
)

type RandomStrategy struct{}

func (s *RandomStrategy) Name() string {
	return "random"
}

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
