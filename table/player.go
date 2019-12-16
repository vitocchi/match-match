package table

import (
	"math"
	"math/rand"

	"github.com/vitocchi/match-match/table/card"
)

// Strategy defines players strategy how decide the card try to pick.
type Strategy interface {
	// DecideFirstTarget receives cardMap value and currentTurn. It decides which card try to pick at first.
	DecideFirstTarget(
		cm card.CardMap,
		currentTurn card.Turn,
	) card.Card
	// DecideSecondTarget receives cardMap value, currentTurn and first picked card. It decides which card try to pick next.
	// cardMap must not include first picked card. so before call this function.
	// So you have to drop the first picked card from cardMap.
	DecideSecondTarget(
		cm card.CardMap,
		currentTurn card.Turn,
		firstPicked card.Card,
	) card.Card
}

// Player is one person in the game.
// Player has a strategy to decide which card try to pick in each situation.
// Player has theirs game point.
// Each Player have to identified by unique name in Players struct.
type Player struct {
	strategy Strategy
	point    uint
	name     string
}

// NewPlayer is constructor of Player
func NewPlayer(s Strategy, n string) Player {
	return Player{
		strategy: s,
		name:     n,
	}
}

// pickCards receives cardMap value which lives in only this function, and currentTurn.
func (p *Player) pickCards(cm card.CardMap, currentTurn card.Turn) [2]card.Card {
	firstTarget := p.strategy.DecideFirstTarget(cm, currentTurn)
	firstPicked := pickCard(firstTarget, cm, currentTurn)
	cm.Drop(firstPicked)
	secondTarget := p.strategy.DecideSecondTarget(cm, currentTurn, firstPicked)
	secondPicked := pickCard(secondTarget, cm, currentTurn)
	return [2]card.Card{firstPicked, secondPicked}
}

// pickCard decides probabilistically which card picked.
func pickCard(target card.Card, cm card.CardMap, currentTurn card.Turn) card.Card {
	probability := calcCardsProbability(uint(len(cm)), int(currentTurn)-int(cm[target]))
	return pickCardsProbabilistically(cm, target, probability)
}

// 0 ~ 1
type probability float64

// cardNum is number of all card left in table.
// elapsedTurn is turn passed after the card have been flipped.
// elapsedTurn of the card which was flipped at last turn is 1.
// elapsedTurn of the card which never been flipped is 0 or less.
func calcCardsProbability(cardNum uint, elapsedTurn int) probability {
	fcn := float64(cardNum)

	// card never been flipped
	if elapsedTurn <= 0 {
		return probability(1 / fcn)
	}

	// t <= 3 + (10/n)
	if float64(elapsedTurn) <= 3+(10/fcn) {
		return 1
	}

	// 0.8^t
	p := math.Pow(float64(0.8), float64(elapsedTurn))

	// 0.8^t <= 1/n
	if p <= 1.0/fcn {
		p = 1.0 / fcn
	}
	return probability(p)
}

func pickCardsProbabilistically(cm card.CardMap, t card.Card, p probability) card.Card {
	// if successed to pick target
	if float64(p) > rand.Float64() {
		return t
	}

	// if failured to pick target
	// pick other card randomly
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
