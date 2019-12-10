package strategy

import (
	"math/rand"

	"github.com/vitocchi/match-match/table/card"
)

type RandomStrategy struct{}

func (s *RandomStrategy) PickCards(cm card.CardMap) [2]card.Card {
	copy := cm.Copy()
	first := s.pickCard(cm)
	copy.Drop(first)
	second := s.pickCard(copy)
	return [2]card.Card{first, second}
}

func (s *RandomStrategy) pickCard(cm card.CardMap) card.Card {
	l := len(cm)
	i := 0

	index := rand.Intn(l)
	for c := range cm {
		if index == i {
			return c
		} else {
			i++
		}
	}
	panic("coudnt pick card")
}
