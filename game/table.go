package game

import (
	"fmt"
)

type Table struct {
	players   []Player
	community []Card
	button    Player
	pot       int
}

func Start() {
	players := []Player{{name: "Daniel", chips: 100}, {name: "Computer", chips: 100}}
	community := []Card{{suit: "♠️", value: "3"}, {suit: "❤️", value: "5"}, {suit: "♣️", value: "5"}}
	var game = Table{players: players, community: community, button: players[0], pot: 13}
	game.Status()
	//deck := GetDeck()
	//deck = deck.Shuffle()
	//fmt.Print(deck)
}

func (g Table) Status() {
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
