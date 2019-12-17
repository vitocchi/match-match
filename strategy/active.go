package strategy

import (
	"github.com/vitocchi/match-match/table/card"
	"fmt"
)

type ActiveStrategy struct{
	TurnThreshold int
}

type pickedMemory struct {
	number int
	lastPickedCard card.Card
	lastPickedAt card.Turn
	secondPickedCard card.Card
	secondPickedAt card.Turn
}

func (p *pickedMemory) turnSum() uint {
	return uint(p.lastPickedAt + p.secondPickedAt)
}

func (p *pickedMemory) haveTwoMemory() bool {
 return p.secondPickedAt != 0 && p.secondPickedCard != card.Card{}
}

type pickedMemoryList []pickedMemory

func newPickedMemoryList(cm card.CardMap) pickedMemoryList {
	list := make(pickedMemoryList, 0, 13)
	for card, turn := range cm {
		list.updatePicked(card, turn)
	}
	return list
}

func (l *pickedMemoryList) updatePicked(card card.Card, pickedAt card.Turn) {
	i, ok := l.findIndex(card)
	if ok {
		if (*l)[i].lastPickedAt < pickedAt {
			(*l)[i].secondPickedCard = (*l)[i].lastPickedCard
			(*l)[i].secondPickedAt = (*l)[i].lastPickedAt
			(*l)[i].lastPickedCard = card
			(*l)[i].lastPickedAt = pickedAt
		} else if (*l)[i].secondPickedAt < pickedAt {
			(*l)[i].secondPickedCard = card
			(*l)[i].secondPickedAt = pickedAt
		}
	} else {
		*l = append(*l, pickedMemory{
			number: card.Number(),
			lastPickedCard: card,
			lastPickedAt: pickedAt,
		})
	}
	
}

func (l *pickedMemoryList) findIndex(card card.Card) (int, bool) {
	for index, entry := range *l {
		if entry.number == card.Number() {
			return index, true
		}
	}
	return 0, false
}

func (l *pickedMemoryList) getMostLikelyMemory() pickedMemory {
	var index int
	var maxSum uint
	for i, m := range *l {
		if maxSum < m.turnSum() {
			maxSum = m.turnSum()
			index = i
		}
	}
	return (*l)[index]
}

func (s *ActiveStrategy) DecideFirstTarget(cm card.CardMap, currentTurn card.Turn) card.Card {
	list := newPickedMemoryList(cm)
	mostlikely := list.getMostLikelyMemory()
	if  int(2*currentTurn) - int(mostlikely.turnSum()) <= s.TurnThreshold  && mostlikely.haveTwoMemory() {
		return  mostlikely.secondPickedCard
	}
	return s.getMostNotLikelyCard(cm)
}

func (s *ActiveStrategy) getMostNotLikelyCard(cm card.CardMap) card.Card {
	var min  uint
	var card card.Card
	for c, turn := range cm {
		if min == 0 || min > uint(turn) {
			card = c
			min = uint(turn)
		}
	}
	fmt.Println("most not likely card")
	fmt.Println(card.String())
	return card
}

func (s *ActiveStrategy) DecideSecondTarget(cm card.CardMap, currentTurn card.Turn, firstPicked card.Card) card.Card {
	for card := range cm {
		if card.IsPair(firstPicked) {
			return card
		}
	}
	target := s.getMostLikelyCardOfNumber(cm, firstPicked.Number())
	if int(currentTurn) - int(cm[target]) < s.TurnThreshold {
		return target
	}
	return s.getMostNotLikelyCard(cm)
}

func (s *ActiveStrategy) getMostLikelyCardOfNumber(cm card.CardMap, number int) card.Card {
	var max  uint
	var card card.Card
	for c, turn := range cm {
		if max < uint(turn) && c.Number() == number {
			card = c
			max = uint(turn)
		}
	}
	return card
}