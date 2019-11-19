package card

import (
	"errors"
	"fmt"
)

type Suit int

const (
	Spade Suit = iota
	Heart
	Diamond
	Club
)

func (s Suit) String() string {
	switch s {
	case Spade:
		return "Spade"
	case Heart:
		return "Heart"
	case Diamond:
		return "Diamond"
	case Club:
		return "Club"
	}
	return ""
}

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

func (c Card) String() string {
	return fmt.Sprintf("suit:%s, number:%d", c.suit, c.number)
}
