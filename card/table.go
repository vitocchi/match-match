package card

import (
	"fmt"
	"math/rand"
)

type Table struct {
	cards  Cards
	player Player
	turn   Turn
}

type Player struct {
	strategy Strategy
	point    uint
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

func (p *Player) getPoint() {
	p.point++
}

func NewTable() Table {
	return Table{
		cards:  NewCards(),
		player: Player{},
	}
}

func (t *Table) ExecGame() {
	for t.isGameGoing() {
		t.execOneTurn()
	}
	fmt.Println("point:", t.player.point)
}

func (t *Table) execOneTurn() {
	cs := t.player.pickCards(t.cards)
	fmt.Println()
	fmt.Println("turn", t.turn)
	fmt.Println(cs[0])
	fmt.Println(cs[1])
	if cs[0].isPair(&cs[1]) {
		t.handleMatch(cs)
	} else {
		t.handleUnmatch(cs)
	}
	t.proceedTurn()
}

func (t *Table) handleMatch(cs [2]Card) {
	fmt.Println("match!!")
	t.cards = t.cards.drop(cs[0])
	t.cards = t.cards.drop(cs[1])
	t.player.getPoint()
}

func (t *Table) handleUnmatch(cs [2]Card) {
	t.cards = t.cards.flip(cs[0], t.turn)
	t.cards = t.cards.flip(cs[1], t.turn)
}

func (t *Table) isGameGoing() bool {
	return len(t.cards) != 0
}

func (t *Table) proceedTurn() {
	t.turn++
}

func (t Table) String() string {
	str := fmt.Sprintln("cards:")
	return str + fmt.Sprintln(t.cards)
}
