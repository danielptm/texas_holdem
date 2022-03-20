package game

import (
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

	table := Table{deck: deck, players: []Player{p, c}}

	var t1 = DealPlayers(table)

	assert.Equal(t, 2, len(t1.players[0].hand))
	assert.Equal(t, 2, len(t1.players[1].hand))
	assert.Equal(t, 48, len(t1.deck))
}

func TestDoFlop(t *testing.T) {
	d := GetDeck()

	var table = Table{
		d, []Player{}, make([]Card, 5), Player{}, 1,
	}

	r1 := DoFlop(table)
	//fmt.Println(r2)

	assert.NotEqual(t, "", r1.community[0])
	assert.NotEqual(t, "", r1.community[1])
	assert.NotEqual(t, "", r1.community[2])
	assert.Equal(t, 49, len(r1.deck))
}

func TestDoTurn(t *testing.T) {
	d := GetDeck()

	var table = Table{
		d, []Player{}, make([]Card, 5), Player{}, 1,
	}

	r1 := DoTurn(table)

	assert.NotEqual(t, "", r1.community[3].value)
	assert.Equal(t, 51, len(r1.deck))
}

func TestDoRiver(t *testing.T) {
	d := GetDeck()
	var table = Table{
		d, []Player{}, make([]Card, 5), Player{}, 1,
	}

	r1 := DoRiver(table)

	assert.NotEqual(t, "", r1.community[4].value)
	assert.Equal(t, 51, len(r1.deck))
}

func TestGetHighCard(t *testing.T) {
	d := GetDeck()
	var table = Table{
		d, []Player{}, []Card{{suit: SPADES, order: 5}, {suit: DIAMONDS, order: 8}, {suit: HEARTS, order: 10}, {suit: CLUBS, order: 11}, {suit: SPADES, order: 6}}, Player{}, 1,
	}

	player := Player{
		name:  "Daniel",
		chips: 1,
		hand:  []Card{{suit: SPADES, order: 3}, {suit: SPADES, order: 5}},
	}

	res := GetHighCardForPlayer(table, player)
	assert.Equal(t, 11, res.order)
}

func TestHasStraightFlush(t *testing.T) {
	table := Table{
		players:   []Player{{name: "Daniel", chips: 100}, {name: "Computer", chips: 100}},
		community: []Card{{suit: HEARTS, value: TWO, order: 2}, {suit: DIAMONDS, value: THREE, order: 3}, {suit: SPADES, value: QUEEN, order: 12}, {suit: SPADES, value: KING, order: 13}, {suit: SPADES, value: ACE, order: 14}},
		button:    Player{name: "Daniel", chips: 100},
		pot:       25,
	}

	p := Player{name: "d", chips: 1, hand: []Card{{suit: SPADES, value: TEN, order: 10}, {suit: SPADES, value: JACK, order: 11}}}

	res1 := GetHighCardForPlayer(table, p)

	assert.Equal(t, 14, res1.order)
}

func TestHasPair(t *testing.T) {
	table := Table{
		players:   []Player{{name: "Daniel", chips: 100}, {name: "Computer", chips: 100}},
		community: []Card{{suit: HEARTS, value: TEN, order: 2}, {suit: DIAMONDS, value: THREE, order: 3}, {suit: SPADES, value: QUEEN, order: 12}, {suit: SPADES, value: KING, order: 13}, {suit: SPADES, value: ACE, order: 14}},
		button:    Player{name: "Daniel", chips: 100},
		pot:       25,
	}

	p := Player{name: "d", chips: 1, hand: []Card{{suit: SPADES, value: TEN, order: 10}, {suit: SPADES, value: JACK, order: 11}}}

	res1 := HasPair(p, table)

	assert.Equal(t, true, res1)
}

func TestHasTwoPair(t *testing.T) {
	table := Table{
		players:   []Player{{name: "Daniel", chips: 100}, {name: "Computer", chips: 100}},
		community: []Card{{suit: HEARTS, value: TEN, order: 2}, {suit: DIAMONDS, value: THREE, order: 3}, {suit: SPADES, value: QUEEN, order: 12}, {suit: SPADES, value: KING, order: 13}, {suit: SPADES, value: ACE, order: 14}},
		button:    Player{name: "Daniel", chips: 100},
		pot:       25,
	}

	p := Player{name: "d", chips: 1, hand: []Card{{suit: SPADES, value: TEN, order: 10}, {suit: SPADES, value: THREE, order: 3}}}

	res1 := HasTwoPair(p, table)

	assert.Equal(t, true, res1)
}

func TestHasThreeOfKind(t *testing.T) {
	table := Table{
		players:   []Player{{name: "Daniel", chips: 100}, {name: "Computer", chips: 100}},
		community: []Card{{suit: HEARTS, value: TEN, order: 2}, {suit: DIAMONDS, value: THREE, order: 3}, {suit: SPADES, value: QUEEN, order: 12}, {suit: SPADES, value: KING, order: 13}, {suit: SPADES, value: ACE, order: 14}},
		button:    Player{name: "Daniel", chips: 100},
		pot:       25,
	}

	p := Player{name: "d", chips: 1, hand: []Card{{suit: SPADES, value: TEN, order: 10}, {suit: CLUBS, value: THREE, order: 3}}}

	res1 := HasTwoPair(p, table)

	assert.Equal(t, true, res1)
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
