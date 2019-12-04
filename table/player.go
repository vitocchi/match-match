package table

import (
	"errors"

	"github.com/vitocchi/match-match/table/card"
)

type Strategy interface {
	PickCards(cs card.Cards) [2]card.Card
}

type Player struct {
	strategy Strategy
	point    uint
	name     string
}

func NewPlayer(s Strategy, n string) Player {
	return Player{
		strategy: s,
		name:     n,
	}
}
func (p *Player) pickCards(cs card.Cards) [2]card.Card {
	return p.strategy.PickCards(cs)
}

func (p *Player) getPoint() {
	p.point++
}

type Players []Player

func (ps *Players) AddPlayer(p Player) error {
	if ps.isNameExist(p.name) {
		return errors.New("name is already exist")
	}
	*ps = append(*ps, p)
	return nil
}

func (ps *Players) isNameExist(n string) bool {
	for _, p := range *ps {
		if p.name == n {
			return true
		}
	}
	return false
}
