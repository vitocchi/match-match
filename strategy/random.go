package strategy

import (
	"math/rand"

	"github.com/vitocchi/match-match/table/card"
)

type RandomStrategy struct{}

func (s *RandomStrategy) DecideFirstTarget(cm card.CardMap, currentTurn card.Turn) card.Card {
	return s.decideTargetRandomly(cm)
}

func (s *RandomStrategy) DecideSecondTarget(cm card.CardMap, currentTurn card.Turn, firstPicked card.Card) card.Card {
	return s.decideTargetRandomly(cm)
}

func (s *RandomStrategy) decideTargetRandomly(cm card.CardMap) card.Card {
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
