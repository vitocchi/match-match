package strategy

import (
	"fmt"

	"github.com/vitocchi/match-match/table/card"
)

type SteadyStratgy struct {
	TurnThreshold int
}

func (s *SteadyStratgy) DecideFirstTarget(cm card.CardMap, currentTurn card.Turn) card.Card {
	list := newPickedMemoryList(cm)
	mostlikely := list.getMostLikelyMemory()
	fmt.Printf("most likely memory:%s\n", mostlikely.String())
	if s.willTryToPickByMemory(mostlikely, currentTurn) {
		fmt.Println("try to pick")
		return mostlikely.secondPickedCard
	}
	fmt.Println("try another one")
	return cm.GetMostRecentlyFlipped()
}

func (s *SteadyStratgy) willTryToPickByMemory(pm pickedMemory, currentTurn card.Turn) bool {
	if !pm.haveTwoMemory() {
		return false
	}
	if int(currentTurn-pm.lastPickedAt) <= s.TurnThreshold && int(currentTurn-pm.secondPickedAt) <= s.TurnThreshold {
		return true
	}
	return false
}

func (s *SteadyStratgy) DecideSecondTarget(cm card.CardMap, currentTurn card.Turn, firstPicked card.Card) card.Card {
	target := cm.GetMostRecentlyFlippedInNumber(firstPicked.Number())
	fmt.Printf("most likey card in number:%v picked at %d\n", target, cm[target])
	if s.willTryToPick(cm[target], currentTurn) {
		fmt.Println("try to pick")
		return target
	}
	fmt.Println("try another one")
	return cm.GetMostRecentlyFlipped()
}

func (s *SteadyStratgy) willTryToPick(pickedAt, currentTurn card.Turn) bool {
	if pickedAt == 0 {
		return false
	}
	if int(currentTurn-pickedAt) <= s.TurnThreshold {
		return true
	}
	return false
}
