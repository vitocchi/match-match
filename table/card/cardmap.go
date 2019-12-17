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
	var err error
	var c Card
	for i := 1; i <= 13; i++ {
		c, err = NewCard(Spade, i)
		if err != nil {
			panic(err)
		}
		cm[c] = 0

		c, err = NewCard(Heart, i)
		if err != nil {
			panic(err)
		}
		cm[c] = 0

		c, err = NewCard(Club, i)
		if err != nil {
			panic(err)
		}
		cm[c] = 0

		c, err = NewCard(Diamond, i)
		if err != nil {
			panic(err)
		}
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
