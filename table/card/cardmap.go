package card

import (
	"fmt"
)

// CardMap is collection of left card and each cards last flipped turn
//
// CAUTION
// When you pass CardMap object to any other function, note that the function can mutate the passed objects value.
// after the function, the cardMap you passed might have lost entries, or have changed its turns.
// If you want not to let the function to change value of CardMap, you have to pass COPY of the object
// like this cm.Copy()
type CardMap map[Card]Turn

// NewCardMap is constructor of Map
func NewCardMap() CardMap {
	cm := CardMap{}
	var c Card
	for i := 1; i <= 13; i++ {
		c = NewCard(Spade, i)
		cm[c] = 0

		c = NewCard(Heart, i)
		cm[c] = 0

		c = NewCard(Club, i)
		cm[c] = 0

		c = NewCard(Diamond, i)
		cm[c] = 0
	}

	return cm
}

// Drop delete card from CardMap
func (cm CardMap) Drop(c Card) {
	delete(cm, c)
}

// Flip update cards last flipped turn
func (cm CardMap) Flip(c Card, t Turn) {
	_, ok := cm[c]
	if !ok {
		fmt.Println(cm)
		fmt.Println(c)
		panic("CardNotFound")
	}
	cm[c] = t
}

// Copy returns copy object of cardMap
func (cm CardMap) Copy() CardMap {
	copy := make(CardMap)
	for key, value := range cm {
		copy[key] = value
	}
	return copy
}

// returns most formaly flipped card.
func (cm CardMap) GetMostFormalyFlipped() Card {
	var min uint
	var card Card
	for c, turn := range cm {
		if (card == Card{}) || min > uint(turn) {
			card = c
			min = uint(turn)
		}
	}
	return card
}

// return most recently flipped card
func (cm CardMap) GetMostRecentlyFlipped() Card {
	var max uint
	var card Card
	for c, turn := range cm {
		if max < uint(turn) || (card == Card{}) {
			card = c
			max = uint(turn)
		}
	}
	return card
}

// return most recently flipped card in number
func (cm CardMap) GetMostRecentlyFlippedInNumber(number int) Card {
	var max uint
	var card Card
	for c, turn := range cm {
		if (max < uint(turn) || (card == Card{})) && c.Number() == number {
			card = c
			max = uint(turn)
		}
	}
	return card
}
