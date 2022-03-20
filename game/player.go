package game

import "fmt"

type Player struct {
	name  string
	chips int
	hand  []Card
}

func (p Player) AddChips(toAdd int) {
	p.chips = p.chips + toAdd
}

func (p Player) SubtractChips(toSubtract int) {
	p.chips = p.chips - toSubtract
}

func (p Player) LookAtHand() string {
	res := fmt.Sprintf("%s, %s", p.hand[0].value+" "+p.hand[0].suit, p.hand[1].value+" "+p.hand[1].suit)
	return res
}
