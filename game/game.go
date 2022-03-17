package game

import (
	"fmt"
)

type Game struct {
	players   []Player
	community []Card
	button    Player
	pot       int
}

func Start() {
	players := []Player{{name: "Daniel", chips: 100}, {name: "Computer", chips: 100}}
	community := []Card{{suit: "spades", value: "3"}, {suit: "hearts", value: "5"}, {suit: "spades", value: "5"}}
	var game = Game{players: players, community: community, button: players[0], pot: 13}
	game.Status()
	//deck := GetDeck()
	//deck = deck.Shuffle()
	//fmt.Print(deck)
}

func (g Game) Status() {
	fmt.Println("")
	fmt.Println("------------------STATE----------------------")
	fmt.Printf("Player: %s, Chips: %d", g.players[0].name, g.players[0].chips)
	fmt.Println("")
	fmt.Printf("Player: %s, Chips: %d", g.players[1].name, g.players[1].chips)
	fmt.Println("")
	fmt.Printf("Community: %s", g.community)
	fmt.Println("")
	fmt.Printf("Button: %s", g.button.name)
	fmt.Println("")
	fmt.Printf("Pot: %d", g.pot)
	fmt.Println("")
	fmt.Println("---------------------------------------------")
	fmt.Println("")
}
