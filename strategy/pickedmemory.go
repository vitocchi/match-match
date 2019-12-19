package strategy

import (
	"fmt"

	"github.com/vitocchi/match-match/table/card"
)

type pickedMemory struct {
	number           int
	lastPickedCard   card.Card
	lastPickedAt     card.Turn
	secondPickedCard card.Card
	secondPickedAt   card.Turn
}

func (p *pickedMemory) turnSum() uint {
	return uint(p.lastPickedAt + p.secondPickedAt)
}

func (p *pickedMemory) haveTwoMemory() bool {
	return p.secondPickedAt != 0 && p.secondPickedCard != card.Card{}
}

func (p *pickedMemory) String() string {
	return fmt.Sprintf("last %v piked at%v\nsecond %v picked at%v", p.lastPickedCard, p.lastPickedAt, p.secondPickedCard, p.secondPickedAt)
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
			number:         card.Number(),
			lastPickedCard: card,
			lastPickedAt:   pickedAt,
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
		if maxSum < m.turnSum() && m.haveTwoMemory() {
			maxSum = m.turnSum()
			index = i
		}
	}
	return (*l)[index]
}
