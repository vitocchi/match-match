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