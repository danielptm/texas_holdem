package game

import (
	"math/rand"
	"time"
)

type Card struct {
	value string
	suit  string
	order int
}

type Deck []Card

func GetDeck() Deck {
	var d Deck

	var suits = []string{"❤️", "♣️", "♠️", "♦️"}
	var values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	var order = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for _, s := range suits {
		for i, v := range values {
			var c = Card{suit: s, value: v, order: order[i]}
			d = append(d, c)
		}
	}
	return d
}

func (d Deck) Shuffle() Deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		randIndex := r.Intn(51)
		d[i], d[randIndex] = d[randIndex], d[i]
	}
	return d
}

func DrawTopCard(d Deck) (Card, Deck) {
	card := d[0]
	d = d[1:]
	return card, d
}
