package table

import (
	"fmt"

	"github.com/vitocchi/match-match/table/card"
)

type Table struct {
	cards              card.Cards
	players            []Player
	startPlayerIndex   uint
	currentPlayerIndex uint
	turn               card.Turn
}

// TODO Playersは外部から注入
func NewTable(p []Player) Table {
	return Table{
		cards:   card.NewCards(),
		players: p,
	}
}

func (t *Table) ResetTable() {
	t.startPlayerIndex = (t.startPlayerIndex + uint(1)) % uint(len(t.players))
	t.cards = card.NewCards()
	t.resetPlayersPoint()
	t.resetTurn()
}

func (t *Table) ExecGame() Result {
	t.setStartPlayerAsCurrent()
	for t.isGameGoing() {
		t.execOneTurn()
	}
	return t.genResult()
}

func (t *Table) setStartPlayerAsCurrent() {
	t.currentPlayerIndex = t.startPlayerIndex
}

func (t *Table) resetPlayersPoint() {
	for i := range t.players {
		t.players[i].point = 0
	}
}

func (t *Table) execOneTurn() {
	cs := t.currentPlayer().pickCards(t.cards)
	if cs[0].IsPair(&cs[1]) {
		t.handleMatch(cs)
	} else {
		t.handleUnmatch(cs)
	}
	t.proceedTurn()
}

func (t *Table) handleMatch(cs [2]card.Card) {
	t.cards = t.cards.Drop(cs[0])
	t.cards = t.cards.Drop(cs[1])
	t.currentPlayer().getPoint()
}

func (t *Table) handleUnmatch(cs [2]card.Card) {
	t.cards = t.cards.Flip(cs[0], t.turn)
	t.cards = t.cards.Flip(cs[1], t.turn)
	t.changePlayer()
}

func (t *Table) isGameGoing() bool {
	return len(t.cards) != 0
}

func (t *Table) proceedTurn() {
	t.turn.Proceed()
}

func (t *Table) resetTurn() {
	t.turn.Reset()
}

func (t *Table) currentPlayer() *Player {
	return &(t.players[t.currentPlayerIndex])
}

func (t *Table) changePlayer() {
	t.currentPlayerIndex = (t.currentPlayerIndex + uint(1)) % uint(len(t.players))
}

func (t *Table) genResult() Result {
	pr := make([]PlayerResult, 0, len(t.players))
	for _, p := range t.players {
		pr = append(pr, NewPlayerResult(p.strategy.Name(), p.point))
	}
	return NewResult(t.turn, pr)
}

func (t Table) String() string {
	str := fmt.Sprintln("cards:")
	return str + fmt.Sprintln(t.cards)
}
