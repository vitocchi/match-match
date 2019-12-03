package card

import "github.com/vitocchi/match-match/table/card"

type Strategy interface {
	pickCards(cs card.Cards) [2]card.Card
}

type Player struct {
	strategy Strategy
	point    uint
}

func NewPlayer(s Strategy) Player {
	return Player{
		strategy: s,
	}
}
func (p *Player) pickCards(cs card.Cards) [2]card.Card {
	return p.strategy.pickCards(cs)
}

func (p *Player) getPoint() {
	p.point++
}
