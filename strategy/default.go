package strategy

import (
	"github.com/vitocchi/match-match/table/card"
)

type DefaultStrategy struct{}

// 0 ~ 100 %
type probability uint

type probabilityTable struct {
	c card.Card
	p probability
}
