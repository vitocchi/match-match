package card

const InitialTurn = 0

type Turn uint

func (t *Turn) Proceed() {
	*t++
}

func (t *Turn) Reset() {
	*t = InitialTurn
}
