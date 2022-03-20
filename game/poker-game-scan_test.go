package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockImpl struct {
	reader ReadPokerinput
}

type MockImpl2 struct {
	reader ReadPokerinput
}

func (MockImpl) getInput() (string, error) {
	return "6", nil
}

func (MockImpl2) getInput() (string, error) {
	return "12", nil
}

//func (MockPokerGameScan)
//
func TestGetChips(t *testing.T) {
	pg := PokerIO{reader: new(MockImpl)}
	res, _ := pg.GetChips()
	assert.Equal(t, 6, res)
}

func TestGetChips2(t *testing.T) {
	pg := PokerIO{reader: new(MockImpl2)}
	res, _ := pg.GetChips()
	assert.Equal(t, -1, res)
}
