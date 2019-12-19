package table

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/vitocchi/match-match/table/card"
)

// Strategy defines players strategy how decide the card try to pick.
type Strategy interface {
	// DecideTargets receives cardMap value and currentTurn. It decides which card try to pick.
	DecideFirstTarget(
		cm card.CardMap,
		currentTurn card.Turn,
	) card.Card
	// DecideSecondTarget receives cardMap value, currentTurn and first picked card. It decides which card try to pick next
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
	strategy   Strategy
	memoryRate float64
	point      uint
	name       string
}

// NewPlayer is constructor of Player
func NewPlayer(s Strategy, name string, memoryRate float64) Player {
	return Player{
		strategy:   s,
		memoryRate: memoryRate,
		name:       name,
	}
}

// pickCards receives cardMap value which lives in only this function, and currentTurn.
func (p *Player) pickCards(cm card.CardMap, currentTurn card.Turn) [2]card.Card {
	fmt.Println("-------------------")
	fmt.Printf("turn:%v\n", currentTurn)
	fmt.Printf("player:%v point:%v\n", p.name, p.point)
	fmt.Printf("map:%v\n", cm)

	firstTarget := p.strategy.DecideFirstTarget(cm.Copy(), currentTurn)
	firstPicked := p.pickCard(firstTarget, cm.Copy(), currentTurn)
	fmt.Printf("firstTarget:%v\n", firstTarget)
	fmt.Printf("firstPicked:%v\n", firstPicked)

	cm.Drop(firstPicked)

	secondTarget := p.strategy.DecideSecondTarget(cm.Copy(), currentTurn, firstPicked)
	secondPicked := p.pickCard(secondTarget, cm.Copy(), currentTurn)
	fmt.Printf("secondTarget:%v\n", secondTarget)
	fmt.Printf("secondPicked:%v\n", secondPicked)

	return [2]card.Card{firstPicked, secondPicked}
}

// pickCard decides probabilistically which card picked.
// it start from trying to pick most likely to success to pick card, and loop until the target
func (p *Player) pickCard(target card.Card, cm card.CardMap, currentTurn card.Turn) card.Card {
	list := newSortedCardList(cm)
	fmt.Println("picking card...")
	for _, e := range list {
		probability := p.calcCardsProbability(uint(len(cm)), int(currentTurn), int(cm[e.key]))
		pickedCard := pickCardsProbabilistically(cm, e.key, probability)
		fmt.Printf("%v,%d => %v,%d (%f))\n", e.key, cm[e.key], pickedCard, cm[pickedCard], probability)
		if e.key == target {
			return pickedCard
		}
		cm.Drop(pickedCard)
	}
	panic("coudnt pick card")
}

// 0 ~ 1
type probability float64

// cardNum is number of all card left in table.
// elapsedTurn is turn passed after the card have been flipped.
// elapsedTurn of the card which was flipped at last turn is 1.
// elapsedTurn of the card which never been flipped is 0 or less.
func (p *Player) calcCardsProbability(cardNum uint, currentTurn int, lastFlippedAt int) probability {
	fcn := float64(cardNum)

	// card never been flipped
	if lastFlippedAt == 0 {
		return probability(1 / fcn)
	}

	elapsedTurn := currentTurn - lastFlippedAt

	// t <= 3 + (10/n)
	if float64(elapsedTurn) <= 3+(10/fcn) {
		return 1
	}

	// memoryRate^(t-3-(10/n))
	pr := math.Pow(p.memoryRate, float64(elapsedTurn)-3.0-(10/fcn))

	// memoryRate^t <= 1/n
	if pr <= 1.0/fcn {
		pr = 1.0 / fcn
	}
	return probability(pr)
}

func pickCardsProbabilistically(cm card.CardMap, t card.Card, p probability) card.Card {
	// if target exists in card map and successed to pick target
	_, ok := cm[t]
	if ok && float64(p) > rand.Float64() {
		return t
	}

	// if targe does not exists in card map failured to pick target
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

func newSortedCardList(cm card.CardMap) cardList {
	l := make(cardList, 0, len(cm))
	for k, v := range cm {
		e := entry{k, v}
		l = append(l, e)
	}
	sort.Sort(l)
	return l
}

type entry struct {
	key   card.Card
	value card.Turn
}

type cardList []entry

func (l cardList) Len() int {
	return len(l)
}

func (l cardList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l cardList) Less(i, j int) bool {
	return (l[i].value > l[j].value)
}
