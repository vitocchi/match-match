package card

import "fmt"

const NUMBER_OF_ALL_CARD = 52

type Cards []Card

func NewCards() Cards {
	cs := make(Cards, 0)
	var err error
	var c Card
	for i := 1; i <= 13; i++ {
		c, err = NewCard(Spade, i)
		if err != nil {
			panic(err)
		}
		cs = append(cs, c)

		c, err = NewCard(Heart, i)
		if err != nil {
			panic(err)
		}
		cs = append(cs, c)

		c, err = NewCard(Club, i)
		if err != nil {
			panic(err)
		}
		cs = append(cs, c)

		c, err = NewCard(Diamond, i)
		if err != nil {
			panic(err)
		}
		cs = append(cs, c)
	}

	return cs
}

func (cs Cards) drop(c Card) Cards {
	for key, element := range cs {
		if c.equals(&element) {
			cs = append(cs[:key], cs[key+1:]...)
			//新しいスライスを用意することがポイント
			n := make(Cards, len(cs))
			copy(n, cs)
			return n
		}
	}
	return cs
}

func (cs Cards) String() string {
	var str string
	for _, m := range cs {
		str += fmt.Sprintln(m)
	}
	return str
}
