package card

import (
	"fmt"
	"math/rand"
)

type Table struct {
	cards  Cards
	player Player
}

type Player struct {
	strategy Strategy
}

type Strategy struct{}

func (s *Strategy) pickCards(cs Cards) [2]Card {
	first := rand.Intn(len(cs))
	var second int
	for {
		second = rand.Intn(len(cs))
		if first != second {
			break
		}
	}
	return [2]Card{cs[first], cs[second]}
}

func (p *Player) pickCards(cs Cards) [2]Card {
	return p.strategy.pickCards(cs)
}

func NewTable() Table {
	return Table{
		cards:  NewCards(),
		player: Player{},
	}
}

func (t *Table) ExecGame() {
	for {
		cs := t.player.pickCards(t.cards)
		fmt.Println("picked", cs[0], cs[1])
		if cs[0].isPair(&cs[1]) {
			fmt.Println("drop!!")
			t.cards = t.cards.drop(cs[0])
			t.cards = t.cards.drop(cs[1])
			fmt.Println("left..", len(t.cards)/2)
		}
		if len(t.cards) == 0 {
			break
		}
	}
}

func (t Table) String() string {
	str := fmt.Sprintln("cards:")
	return str + fmt.Sprintln(t.cards)
}
