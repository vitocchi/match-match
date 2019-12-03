package card

import (
	"fmt"

	"github.com/vitocchi/match-match/table/card"
)

type Table struct {
	cards              card.Cards
	players            []Player
	currentPlayerIndex uint
	turn               Turn
}

// TODO Playersは外部から注入
func NewTable(p []Player) Table {
	return Table{
		cards:   card.NewCards(),
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
	fmt.Println("player", t.currentPlayerIndex)
	fmt.Println(cs[0])
	fmt.Println(cs[1])
	if cs[0].IsPair(&cs[1]) {
		t.handleMatch(cs)
	} else {
		t.handleUnmatch(cs)
	}
	t.proceedTurn()
}

func (t *Table) handleMatch(cs [2]card.Card) {
	fmt.Println("match!!")
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
	t.turn++
}

func (t *Table) currentPlayer() *Player {
	return &(t.players[t.currentPlayerIndex])
}

func (t *Table) changePlayer() {
	fmt.Println("player change!!")
	t.currentPlayerIndex = (t.currentPlayerIndex + uint(1)) % uint(len(t.players))
}

func (t Table) String() string {
	str := fmt.Sprintln("cards:")
	return str + fmt.Sprintln(t.cards)
}
