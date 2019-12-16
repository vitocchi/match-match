package table

import (
	"fmt"

	"github.com/vitocchi/match-match/table/card"
)

// Table orchestrates match-match game
type Table struct {
	cardMap            card.CardMap
	players            Players
	startPlayerIndex   uint
	currentPlayerIndex uint
	turn               card.Turn
}

// NewTable is constructor of Table
func NewTable(p Players) Table {
	return Table{
		cardMap: card.NewCardMap(),
		players: p,
	}
}

// Reset delete states of current game and set to next game
func (t *Table) Reset() {
	t.startPlayerIndex = (t.startPlayerIndex + uint(1)) % uint(len(t.players))
	t.cardMap = card.NewCardMap()
	t.resetPlayersPoint()
	t.resetTurn()
}

// ExecGame is simulate one game and return game result
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
	cs := t.currentPlayer().pickCards(t.cardMap.Copy(), t.turn)
	if cs[0].IsPair(cs[1]) {
		t.handleMatch(cs)
	} else {
		t.handleUnmatch(cs)
	}
	t.proceedTurn()
}

func (t *Table) handleMatch(cs [2]card.Card) {
	t.cardMap.Drop(cs[0])
	t.cardMap.Drop(cs[1])
	t.currentPlayer().getPoint()
}

func (t *Table) handleUnmatch(cs [2]card.Card) {
	t.cardMap.Flip(cs[0], t.turn)
	t.cardMap.Flip(cs[1], t.turn)
	t.changePlayer()
}

func (t *Table) isGameGoing() bool {
	return len(t.cardMap) != 0
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
	r := Result{}
	for _, p := range t.players {
		r[p.name] = p.point
	}
	return r
}

func (t Table) String() string {
	str := fmt.Sprintln("cards:")
	return str + fmt.Sprintln(t.cardMap)
}
