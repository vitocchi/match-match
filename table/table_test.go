package table

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/vitocchi/match-match/table/card"
)

type mockStrategy struct{}

func (m mockStrategy) DecideFirstTarget(cm card.CardMap, currentTurn card.Turn) card.Card {
	return card.Card{}
}
func (m mockStrategy) DecideSecondTarget(cm card.CardMap, currentTurn card.Turn, firstPicked card.Card) card.Card {
	return card.Card{}
}

func TestChangeSheets(t *testing.T) {
	p1 := NewPlayer(mockStrategy{}, "1", 0)
	p2 := NewPlayer(mockStrategy{}, "2", 0)
	p3 := NewPlayer(mockStrategy{}, "3", 0)
	p4 := NewPlayer(mockStrategy{}, "4", 0)
	ps := Players{
		p1, p2, p3, p4,
	}

	expecteds := [13]Players{
		{p1, p2, p3, p4},
		{p2, p1, p3, p4},
		{p2, p3, p1, p4},
		{p2, p3, p4, p1},
		{p1, p3, p4, p2},
		{p3, p1, p4, p2},
		{p3, p4, p1, p2},
		{p3, p4, p2, p1},
		{p1, p4, p2, p3},
		{p4, p1, p2, p3},
		{p4, p2, p1, p3},
		{p4, p2, p3, p1},
		{p1, p2, p3, p4},
	}

	table := NewTable(ps)
	for key, expected := range expecteds {
		if !reflect.DeepEqual(table.players, expected) {
			t.Fatalf("in %d, expected %v actual %v", key, expected, table.players)
		}
		table.changeSeats()
		fmt.Println(table.sheetChanger)
	}
}
