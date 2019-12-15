package strategy

import (
	"math"
	"math/rand"

	"github.com/vitocchi/match-match/table/card"
)

// 0 ~ 1
type probability float64

func calcCardsProbability(cardNum uint, elapsedTurn int) probability {
	fcn := float64(cardNum)
	if elapsedTurn < 0 {
		return probability(1 / fcn)
	}
	if float64(elapsedTurn) <= 3+(10/fcn) {
		return 1
	}

	p := math.Pow(float64(0.8), float64(elapsedTurn))
	if p <= 1.0/fcn {
		p = 1.0 / fcn
	}
	return probability(p)
}

func pickCardsProbabilistically(cm card.CardMap, t card.Card, p probability) card.Card {
	if float64(p) > rand.Float64() {
		return t
	}
	l := len(cm)
	for {
		i := 0
		index := rand.Intn(l)
		for c := range cm {
			if index == i && c != t {
				return c
			}
			i++
		}
	}
}
