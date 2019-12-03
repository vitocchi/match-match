package main

import (
	"fmt"

	"github.com/vitocchi/neurathenia/card"
)

func main() {
	table := card.NewTable()
	table.ExecGame()
	fmt.Println(table)
}
