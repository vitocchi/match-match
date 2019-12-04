package card

const INITIAL_TURN = 0

type Turn uint

func (t *Turn) Proceed() {
	*t++
}

func (t *Turn) Reset() {
	*t = INITIAL_TURN
}
