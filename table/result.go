package table

import (
	"github.com/vitocchi/match-match/table/card"
)

type Result struct {
	Turn          card.Turn      `json:"turn"`
	PlayerResults []PlayerResult `json:"result"`
}

func NewResult(t card.Turn, pr []PlayerResult) Result {
	return Result{
		Turn:          t,
		PlayerResults: pr,
	}
}

type PlayerResult struct {
	Strategy string `json:"strategy"`
	Point    uint   `json:"point"`
}

func NewPlayerResult(strategy string, point uint) PlayerResult {
	return PlayerResult{
		Strategy: strategy,
		Point:    point,
	}
}
