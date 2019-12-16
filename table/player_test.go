package table

import (
	"testing"
)

func TestCalcCardProbability(t *testing.T) {
	// 100 % のパターン
	calcAndAssert(10, 4, 1, t)
	calcAndAssert(10, 3, 1, t)
	calcAndAssert(10, 0, 0.1, t)
	calcAndAssert(100, 0, 0.01, t)
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

func calcAndAssert(cardNum uint, elapsedTrun int, expected probability, t *testing.T) {
	var p probability
	p = calcCardsProbability(cardNum, elapsedTrun)
	if p != expected {
		t.Errorf("cardNum: %d, turn:%d, expected:%f, acutual:%f", cardNum, elapsedTrun, expected, p)
	}
}
