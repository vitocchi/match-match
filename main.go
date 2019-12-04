package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/vitocchi/match-match/strategy"
	"github.com/vitocchi/match-match/table"
)

const SimulationTime = 1000

type SimulationResult []table.Result

func main() {
	rand.Seed(time.Now().UnixNano())
	p := initPlayers()
	result := execSimulation(p)
	fmt.Println(result.toJSON())
}

func initPlayers() []table.Player {
	p := make([]table.Player, 0, 1)
	p = append(p, table.NewPlayer(&strategy.RandomStrategy{}))
	return p
}

func execSimulation(p []table.Player) SimulationResult {
	t := table.NewTable(p)
	results := make([]table.Result, 0, SimulationTime)
	for i := 0; i < SimulationTime; i++ {
		r := t.ExecGame()
		results = append(results, r)
		t.Reset()
	}
	return results
}

func (r *SimulationResult) toJSON() string {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}
