package card

import "fmt"

type memory struct {
	card        Card
	remembrance int
}

func newMemory(c Card) memory {
	return memory{
		card:        c,
		remembrance: 0,
	}
}

func (m memory) String() string {
	return fmt.Sprintf("%s, remembrance=%d", m.card, m.remembrance)
}

type Memories []memory

func NewMemories() Memories {
	var ms Memories
	var c Card
	var err error
	var m memory
	for i := 1; i <= 13; i++ {
		c, err = NewCard(Spade, i)
		if err != nil {
			panic(err)
		}
		m = newMemory(c)
		ms = append(ms, m)

		c, err = NewCard(Heart, i)
		if err != nil {
			panic(err)
		}
		m = newMemory(c)
		ms = append(ms, m)

		c, err = NewCard(Club, i)
		if err != nil {
			panic(err)
		}
		m = newMemory(c)
		ms = append(ms, m)

		c, err = NewCard(Diamond, i)
		if err != nil {
			panic(err)
		}
		m = newMemory(c)
		ms = append(ms, m)
	}

	return ms
}

func (ms Memories) String() string {
	var str string
	for _, m := range ms {
		str += fmt.Sprintln(m)
	}
	return str
}
