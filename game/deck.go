package game

import (
	"math/rand"
	"time"
)

type Card struct {
	suit  string
	value string
}

type deck []Card

func GetDeck() deck {
	var d deck

	var suits = []string{"spades", "clubs", "hearts", "diamonds"}
	var values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for _, s := range suits {
		for _, v := range values {
			var c = Card{suit: s, value: v}
			d = append(d, c)
		}
	}
	return d
}

func (d deck) Shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		randIndex := r.Intn(51)
		d[i], d[randIndex] = d[randIndex], d[i]
	}
	return d
}

func DrawTopCard(d deck) (Card, deck) {
	card := d[0]
	d = d[1:]
	return card, d
}
