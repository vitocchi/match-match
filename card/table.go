package card

import (
	"fmt"
	"math/rand"
)

type Table struct {
	cards              Cards
	players            []Player
	currentPlayerIndex uint
	turn               Turn
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

// TODO Playersは外部から注入
func NewTable() Table {
	p := make([]Player, 4, 4)
	return Table{
		cards:   NewCards(),
		players: p,
	}
}

func (t *Table) ExecGame() {
	for t.isGameGoing() {
		t.execOneTurn()
	}
	fmt.Println(t.players)
}

func (t *Table) execOneTurn() {
	cs := t.currentPlayer().pickCards(t.cards)
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
	t.currentPlayer().getPoint()
}

func (t *Table) handleUnmatch(cs [2]Card) {
	t.cards = t.cards.flip(cs[0], t.turn)
	t.cards = t.cards.flip(cs[1], t.turn)
	t.changePlayer()
}

func (t *Table) isGameGoing() bool {
	return len(t.cards) != 0
}

func (t *Table) proceedTurn() {
	t.turn++
}

func (t *Table) currentPlayer() *Player {
	return &(t.players[t.currentPlayerIndex])
}

func (t *Table) changePlayer() {
	t.currentPlayerIndex = (t.currentPlayerIndex + uint(1)) % uint(len(t.players))
}

func (t Table) String() string {
	str := fmt.Sprintln("cards:")
	return str + fmt.Sprintln(t.cards)
}
