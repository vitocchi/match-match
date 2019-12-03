package card

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
		return "♤"
	case Heart:
		return "♡"
	case Diamond:
		return "♢"
	case Club:
		return "♧"
	}
	return ""
}
