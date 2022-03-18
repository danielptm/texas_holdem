package game

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDealPlayers(t *testing.T) {
	deck := GetDeck()
	p := Player{
		name:  "Daniel",
		chips: 100,
		hand:  make([]Card, 2),
	}
	c := Player{
		name:  "Daniel",
		chips: 100,
		hand:  make([]Card, 2),
	}

	var p1, p2, d = DealPlayers(p, c, deck)

	assert.Equal(t, 2, len(p1.hand))
	assert.Equal(t, 2, len(p2.hand))
	assert.Equal(t, 48, len(d))
}

func TestDoFlop(t *testing.T) {
	var table = Table{
		[]Player{}, make([]Card, 5), Player{}, 1,
	}

	d := GetDeck()

	r1, r2 := DoFlop(table, d)
	fmt.Println(r2)

	assert.NotEqual(t, "", r1.community[0])
	assert.NotEqual(t, "", r1.community[1])
	assert.NotEqual(t, "", r1.community[2])
	assert.Equal(t, 49, len(r2))
}

func TestDoTurn(t *testing.T) {
	var table = Table{
		[]Player{}, make([]Card, 5), Player{}, 1,
	}

	d := GetDeck()

	r1, r2 := DoTurn(table, d)

	assert.NotEqual(t, "", r1.community[3].value)
	assert.Equal(t, 51, len(r2))
}

func TestDoRiver(t *testing.T) {
	var table = Table{
		[]Player{}, make([]Card, 5), Player{}, 1,
	}

	d := GetDeck()

	r1, r2 := DoRiver(table, d)

	assert.NotEqual(t, "", r1.community[4].value)
	assert.Equal(t, 51, len(r2))
}

func TestGetHighCard(t *testing.T) {
	var table = Table{
		[]Player{}, []Card{{suit: SPADES, order: 5}, {suit: DIAMONDS, order: 8}, {suit: HEARTS, order: 10}, {suit: CLUBS, order: 11}, {suit: SPADES, order: 6}}, Player{}, 1,
	}

	player := Player{
		name:  "Daniel",
		chips: 1,
		hand:  []Card{{suit: SPADES, order: 3}, {suit: SPADES, order: 5}},
	}

	res := GetHighCardForPlayer(table, player)
	assert.Equal(t, 11, res.order)
}

//func TestHasRoyalFlush(t *testing.T) {
//	table := Table{
//		players:   []Player{{name: "Daniel", chips: 100}, {name: "Computer", chips: 100}},
//		community: []Card{{suit: HEARTS, value: TWO}, {suit: DIAMONDS, value: THREE}, {suit: SPADES, value: QUEEN}, {suit: SPADES, value: KING}, {suit: SPADES, value: ACE}},
//		button:    Player{name: "Daniel", chips: 100},
//		pot:       25,
//	}
//
//	p := Player{name: "d", chips: 1, hand: []Card{{suit: SPADES, value: TEN}, {suit: SPADES, value: JACK}}}
//
//	res := HasRoyalFlush(table, p)
//
//	assert.Equal(t, true, res)
//}

//func TestHasHighCard(t *testing.T) {
//	table := Table{
//		players:   []Player{{name: "Daniel", chips: 100}, {name: "Computer", chips: 100}},
//		community: []Card{{suit: HEARTS, value: TWO}, {suit: DIAMONDS, value: THREE}, {suit: SPADES, value: QUEEN}, {suit: SPADES, value: KING}, {suit: SPADES, value: ACE}},
//		button:    Player{name: "Daniel", chips: 100},
//		pot:       25,
//	}
//
//	p := Player{name: "d", chips: 1, hand: []Card{{suit: SPADES, value: TEN}, {suit: SPADES, value: JACK}}}
//
//}
