package table

import (
	"math"
	"math/rand"

	"github.com/vitocchi/match-match/table/card"
)

type Strategy interface {
	DecideFirstTarget(cm card.CardMap, currentTurn card.Turn) card.Card
	DecideSecondTarget(cm card.CardMap, currentTurn card.Turn, firstPicked card.Card) card.Card
}

type Player struct {
	strategy Strategy
	point    uint
	name     string
}

func NewPlayer(s Strategy, n string) Player {
	return Player{
		strategy: s,
		name:     n,
	}
}

func (p *Player) pickCards(cm card.CardMap, currentTurn card.Turn) [2]card.Card {
	firstTarget := p.strategy.DecideFirstTarget(cm, currentTurn)
	firstPicked := pickCard(firstTarget, cm, currentTurn)
	cm.Drop(firstPicked)
	secondTarget := p.strategy.DecideSecondTarget(cm, currentTurn, firstPicked)
	secondPicked := pickCard(secondTarget, cm, currentTurn)
	return [2]card.Card{firstPicked, secondPicked}
}

func pickCard(target card.Card, cm card.CardMap, currentTurn card.Turn) card.Card {
	probability := calcCardsProbability(uint(len(cm)), int(currentTurn)-int(cm[target]))
	return pickCardsProbabilistically(cm, target, probability)
}

// 0 ~ 1
type probability float64

func calcCardsProbability(cardNum uint, elapsedTurn int) probability {
	fcn := float64(cardNum)
	if elapsedTurn <= 0 {
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

func (p *Player) getPoint() {
	p.point++
}
