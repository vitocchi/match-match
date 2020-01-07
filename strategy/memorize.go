package strategy

import (
	"fmt"

	"github.com/vitocchi/match-match/table/card"
)

type MemorizeStrategy struct {
	TurnThreshold int
}

func (s *MemorizeStrategy) DecideFirstTarget(cm card.CardMap, currentTurn card.Turn) card.Card {
	list := newPickedMemoryList(cm)
	mostlikely := list.getMostLikelyMemory()
	fmt.Printf("most likely memory:%s\n", mostlikely.String())
	if s.willTryToPickByMemory(mostlikely, currentTurn, len(cm)) {
		fmt.Println("try to pick")
		return mostlikely.secondPickedCard
	}
	fmt.Println("try another one")
	return cm.GetMostFormalyFlipped()
}

// returns if try to pick memories card.
// if memory doesnt have two cards memory, it returns false
// if sum of elapsedTurn of cards memory is more than threshold, it return false
func (s *MemorizeStrategy) willTryToPickByMemory(pm pickedMemory, currentTurn card.Turn, cardNum int) bool {
	if !pm.haveTwoMemory() {
		return false
	}
	if int(2*currentTurn)-int(pm.turnSum()) > s.TurnThreshold+52/cardNum {
		return false
	}
	return true
}

func (s *MemorizeStrategy) DecideSecondTarget(cm card.CardMap, currentTurn card.Turn, firstPicked card.Card) card.Card {
	target := cm.GetMostRecentlyFlippedInNumber(firstPicked.Number())
	fmt.Printf("most likey card in number:%v picked at %d\n", target, cm[target])
	if s.willTryToPick(cm[target], currentTurn, len(cm)) {
		fmt.Println("try to pick")
		return target
	}
	fmt.Println("try another one")
	return cm.GetMostFormalyFlipped()
}

func (s *MemorizeStrategy) willTryToPick(pickedAt, currentTurn card.Turn, cardNum int) bool {
	if pickedAt == 0 {
		return false
	}
	return int(currentTurn)-int(pickedAt) < s.TurnThreshold/2+52/cardNum
}
