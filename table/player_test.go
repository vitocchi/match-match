package table

import (
	"math"
	"testing"
)

func TestCalcCardProbability(t *testing.T) {
	// 100 % のパターン
	calcAndAssert(10, 4, 4, 1, t)
	calcAndAssert(10, 5, 2, 1, t)
	calcAndAssert(10, 1, 0, 0.1, t)
	calcAndAssert(100, 100000000, 20, 0.01, t)

	calcAndAssert(10, 5, 1, 1, t)
	calcAndAssert(10, 6, 1, 0.8, t)
	calcAndAssert(10, 7, 1, probability(math.Pow(0.8, 2)), t)
	calcAndAssert(10, 21, 1, 1.0/10, t)

	calcAndAssert(5, 5, 1, 1, t)
	calcAndAssert(5, 6, 1, 1, t)
	calcAndAssert(5, 7, 1, 0.8, t)
	calcAndAssert(5, 8, 1, probability(math.Pow(0.8, 2)), t)
	calcAndAssert(5, 21, 1, 1.0/5, t)
	/*
		calcAndAssert(, , , t)
		calcAndAssert(, , , t)
		calcAndAssert(, , , t)
		calcAndAssert(, , , t)
		calcAndAssert(, , , t)
		calcAndAssert(, , , t)
		calcAndAssert(, , , t)
		calcAndAssert(, , , t)
		calcAndAssert(, , , t)
	*/
}
func calcAndAssert(cardNum uint, currentTrun int, lastFlippedAt int, expected probability, t *testing.T) {
	var p probability
	player := NewPlayer(mockStrategy{}, "mock", 0.8)
	p = player.calcCardsProbability(cardNum, currentTrun, lastFlippedAt)
	if p != expected {
		t.Errorf("cardNum: %d, currentTurn:%d, lastFlippedAt %d, expected:%f, acutual:%f", cardNum, currentTrun, lastFlippedAt, expected, p)
	}
}
