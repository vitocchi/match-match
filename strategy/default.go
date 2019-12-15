package strategy

import (
	"github.com/vitocchi/match-match/table/card"
)

type DefaultStrategy struct{}

func (s *DefaultStrategy) PickCards(cm card.CardMap, currentTurn card.Turn) [2]card.Card {
	firstTarget := s.decideFirstTarget(cm, currentTurn)
	firstPicked := s.tryToPick(firstTarget, cm, currentTurn)
	cm.Drop(firstPicked)
	secondTarget := s.decideSecondTarget(cm, firstPicked, currentTurn)
	secondPicked := s.tryToPick(secondTarget, cm, currentTurn)
	return [2]card.Card{firstPicked, secondPicked}
}

func (s *DefaultStrategy) decideFirstTarget(cm card.CardMap, currentTurn card.Turn) card.Card {
	var lastTurn card.Turn
	var lastFlipped card.Card
	for card, turn := range cm {
		if lastTurn <= turn {
			lastFlipped = card
			lastTurn = turn
		}
	}
	return lastFlipped
}

func (s *DefaultStrategy) decideSecondTarget(cm card.CardMap, firstPicked card.Card, currentTurn card.Turn) card.Card {
	for card := range cm {
		if card.IsPair(firstPicked) {
			return card
		}
	}
	panic("coudnt decide first target")
}

func (s *DefaultStrategy) tryToPick(target card.Card, cm card.CardMap, currentTurn card.Turn) card.Card {
	p := calcCardsProbability(uint(len(cm)), int(currentTurn)-int(cm[target]))
	return pickCardsProbabilistically(cm, target, p)
}
