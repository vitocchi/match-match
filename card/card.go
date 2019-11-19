package card

import "errors"

type Suit int

const (
	Spade Suit = iota
	Heart
	Diamond
	Club
)

type Card struct {
	suit   Suit
	number int
}

func NewCard(s Suit, n int) (Card, error) {
	if isNumberValid(n) {
		return Card{
			suit:   s,
			number: n,
		}, nil
	}

	return Card{}, errors.New("number is invalid")
}

func isNumberValid(n int) bool {
	return n >= 1 && n <= 13
}
