package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/vitocchi/match-match/strategy"
	"github.com/vitocchi/match-match/table"
)

const SimulationTime = 1200

type SimulationResult []table.Result

func main() {
	initSeed()
	p := initPlayers()
	fmt.Println("------start------------")
	result := execSimulation(p)
	fmt.Println(result.toJSON())

	file := openResultFile()
	result.writeCSV(file)
	file.Close()
}

func initSeed() {
	rand.Seed(time.Now().UnixNano())
}

func initPlayers() []table.Player {
	p := make(table.Players, 0, 4)
	p.AddPlayer(table.NewPlayer(&strategy.RandomStrategy{}, "nobita", 0))
	p.AddPlayer(table.NewPlayer(&strategy.SteadyStratgy{4}, "suneo", 0.8))
	p.AddPlayer(table.NewPlayer(&strategy.MemorizeStrategy{10}, "sizuka", 0.8))
	p.AddPlayer(table.NewPlayer(&strategy.MemorizeStrategy{1000}, "jaian", 0.6))
	return p
}

func execSimulation(p table.Players) SimulationResult {
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

func openResultFile() *os.File {
	name := strconv.FormatInt(time.Now().Unix(), 10) + ".csv"
	file, err := os.Create("result/" + name)
	if err != nil {
		panic(err)
	}
	return file
}

func (r *SimulationResult) writeCSV(w io.Writer) {
	table := r.toCSVTable()
	writer := csv.NewWriter(w)
	writer.WriteAll(table)
	if err := writer.Error(); err != nil {
		panic(err)
	}
}

func (r *SimulationResult) toCSVTable() [][]string {
	if len(*r) == 0 {
		panic("result is emply")
	}
	var table [][]string
	header := make([]string, 0, len((*r)[0]))
	for playerName := range (*r)[0] {
		header = append(header, playerName)
	}
	table = append(table, header)
	for _, result := range *r {
		var record []string
		for _, playerName := range header {
			record = append(record, strconv.Itoa(int(result[playerName])))
		}
		table = append(table, record)
	}
	return table
}
