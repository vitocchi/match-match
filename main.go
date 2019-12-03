package main

import (
	"math/rand"
	"time"

	"github.com/vitocchi/neurathenia/card"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	table := card.NewTable()
	table.ExecGame()
}
