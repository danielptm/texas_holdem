package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockPokerGameScan struct {
	reader ReadPokerinput
}

type MockPokerGameScan2 struct {
	reader ReadPokerinput
}

func (MockPokerGameScan) getInput() (string, error) {
	return "6", nil
}

func (MockPokerGameScan2) getInput() (string, error) {
	return "12", nil
}

//func (MockPokerGameScan)
//
func TestGetChips(t *testing.T) {
	pg := PokerGameScan{reader: new(MockPokerGameScan)}
	res, _ := pg.GetChips()
	assert.Equal(t, 6, res)
}

func TestGetChips2(t *testing.T) {
	pg := PokerGameScan{reader: new(MockPokerGameScan2)}
	res, _ := pg.GetChips()
	assert.Equal(t, -1, res)
}
