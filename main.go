package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/vitocchi/match-match/table"
	"github.com/vitocchi/match-match/table/card"
)

type RandomStrategy struct{}

func (s *RandomStrategy) Name() string {
	return "random"
}

func (s *RandomStrategy) PickCards(cs card.Cards) [2]card.Card {
	first := rand.Intn(len(cs))
	var second int
	for {
		second = rand.Intn(len(cs))
		if first != second {
			break
		}
	}
	return [2]card.Card{cs[first], cs[second]}
}

const SimulationTime = 1000

func main() {
	rand.Seed(time.Now().UnixNano())
	p := make([]table.Player, 1, 1)
	p[0] = table.NewPlayer(&RandomStrategy{})
	t := table.NewTable(p)
	results := make([]table.Result, 0, SimulationTime)

	for i := 0; i < SimulationTime; i++ {
		r := t.ExecGame()
		results = append(results, r)
		t.ResetTable()
	}
	fmt.Println(resultsToJSON(results))
}

func resultsToJSON(r []table.Result) string {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}
