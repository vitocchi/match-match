package card

import (
	"fmt"
)

type Card struct {
	suit   Suit
	number int
}

func NewCard(s Suit, n int) Card {
	if isNumberValid(n) {
		return Card{
			suit:   s,
			number: n,
		}
	}

	panic("number is invalid")
}

func (c *Card) Number() int {
	return c.number
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
