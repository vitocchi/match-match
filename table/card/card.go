package card

import (
	"errors"
	"fmt"
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

func (c *Card) IsPair(other Card) bool {
	return c.suit != other.suit && c.number == other.number
}

func isNumberValid(n int) bool {
	return n >= 1 && n <= 13
}

func (c Card) String() string {
	return fmt.Sprintf("%s %2d", c.suit, c.number)
}
