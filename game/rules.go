package game

/**
Play a hand return true if a next hand should happen, otherwise false if game is over.
*/
func PlayHand(table Table) bool {
	//Deal cards to players
	table.deck = table.deck.Shuffle()
	table = DealPlayers(table)
	println("Welcome to Daniel's GO build texas holdem for 2 players.")
	playerBet, _ := PokerIO{reader: new(ScanImpl)}.GetChips()
	computerBet := 5
	table.players[0].chips = table.players[0].chips - computerBet
	table.players[1].chips = table.players[1].chips - playerBet
	table.pot = table.pot + (playerBet + computerBet)
	table.Status("PRE-FLOP")
	table = DoFlop(table)
	table.Status("POST-FLOP")
	playerBet2, _ := PokerIO{reader: new(ScanImpl)}.GetChips()
	computerBet2 := 5
	table.players[0].chips = table.players[0].chips - computerBet2
	table.players[1].chips = table.players[1].chips - playerBet2
	table.pot = table.pot + (playerBet2 + computerBet2)
	table = DoTurn(table)
	table.Status("POST-TURN")
	playerBet3, _ := PokerIO{reader: new(ScanImpl)}.GetChips()
	computerBet3 := 5
	table.players[0].chips = table.players[0].chips - computerBet3
	table.players[1].chips = table.players[1].chips - playerBet3
	table.pot = table.pot + (playerBet3 + computerBet3)
	table = DoRiver(table)
	table.Status("POST-RIVER")
	playerBet4, _ := PokerIO{reader: new(ScanImpl)}.GetChips()
	computerBet4 := 5
	table.players[0].chips = table.players[0].chips - computerBet4
	table.players[1].chips = table.players[1].chips - playerBet4
	table.pot = table.pot + (playerBet4 + computerBet4)
	return true
}

func DealPlayers(table Table) Table {
	cards1 := table.deck[0:2]
	cards2 := table.deck[2:4]
	table.players[0].hand = cards1
	table.players[1].hand = cards2
	table.deck = table.deck[4:]
	return table
}

func DoFlop(table Table) Table {
	table.community[0], table.community[1], table.community[2] = table.deck[0], table.deck[1], table.deck[2]
	table.deck = table.deck[3:]
	return table
}

func DoTurn(table Table) Table {
	table.community[3] = table.deck[0]
	table.deck = table.deck[1:]
	return table
}

func DoRiver(table Table) Table {
	table.community[4] = table.deck[0]
	table.deck = table.deck[1:]
	return table
}

func GetBestHand(table Table, player Player) string {
	if HasThreeOfKind(player, table) {
		return THREE_OF_A_KIND
	}
	if HasTwoPair(player, table) {
		return TWO_PAIR
	}
	if HasPair(player, table) {
		return PAIR
	}
	return HIGH_CARD
}

func GetHighCardForPlayer(table Table, player Player) Card {
	highCard := Card{
		order: 1,
	}
	for _, c := range table.community {
		if c.order > highCard.order {
			highCard = c
		}
	}

	for _, c := range player.hand {
		if c.order > highCard.order {
			highCard = c
		}
	}
	return highCard
}

func HasPair(player Player, table Table) bool {
	for _, p_c := range player.hand {
		for _, t_c := range table.community {
			if p_c.value == t_c.value {
				return true
			}
		}
	}
	return false
}

func HasTwoPair(player Player, table Table) bool {
	pairNums := 0
	for _, p_c := range player.hand {
		for _, t_c := range table.community {
			if p_c.value == t_c.value {
				pairNums = pairNums + 2
			}
		}
	}
	if pairNums >= 4 {
		return true
	} else {
		return false
	}
}

func HasThreeOfKind(player Player, table Table) bool {
	pairNums := 0
	for _, p_c := range player.hand {
		for _, t_c := range table.community {
			if p_c.value == t_c.value {
				pairNums = pairNums + 1
			}
		}
	}
	if pairNums == 3 {
		return true
	} else {
		return false
	}
}

func HasStraight() bool {
	return false
}

func HasFlush() bool {
	return false
}

//func HasStraightFlush() bool {
//	return ""
//}
//

//
//func HasFourOfKind() bool {
//	return ""
//}
//

//

//func HasRoyalFlush(table Table, player Player) bool {
//
//	return false
//}
