package game

type Player struct {
	name  string
	chips int
}

func (p Player) AddChips(toAdd int) {
	p.chips = p.chips + toAdd
}

func (p Player) SubtractChips(toSubtract int) {
	p.chips = p.chips - toSubtract
}
