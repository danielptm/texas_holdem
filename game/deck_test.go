package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDeck(t *testing.T) {
	d := GetDeck()
	assert.Equal(t, 52, len(d))
}

func TestShuffle(t *testing.T) {
	d1 := GetDeck()
	d2 := GetDeck()
	var res = d1.Shuffle()
	var res2 = d2.Shuffle()
	assert.NotEqualf(t, res, res2, "Shuffle produces the same deck")
}

func TestDeck_DrawTopCard(t *testing.T) {
	var d = GetDeck()
	var _, newDeck = DrawTopCard(d)
	assert.Equal(t, 51, len(newDeck))
}
