package card

import "fmt"

const NUMBER_OF_ALL_CARD = 52

type CardMap map[Card]Turn

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

func (cm CardMap) Drop(c Card) {
	delete(cm, c)
}

func (cm CardMap) Flip(c Card, t Turn) {
	_, ok := cm[c]
	if !ok {
		fmt.Println(cm)
		fmt.Println(c)
		panic("CardNotFound")
	}
	cm[c] = t
}

func (cm CardMap) Copy() CardMap {
	copy := make(CardMap)
	for key, value := range cm {
		copy[key] = value
	}
	return copy
}
