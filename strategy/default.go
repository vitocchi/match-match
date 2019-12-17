package strategy

import "github.com/vitocchi/match-match/table/card"

type DefaultStrategy struct{}

func (s *DefaultStrategy) DecideFirstTarget(cm card.CardMap, currentTurn card.Turn) card.Card {
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

func (s *DefaultStrategy) DecideSecondTarget(cm card.CardMap, currentTurn card.Turn, firstPicked card.Card) card.Card {
	for card := range cm {
		if card.IsPair(firstPicked) {
			return card
		}
	}
	panic("coudnt decide first target")
}
