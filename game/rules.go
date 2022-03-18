package game

func DealPlayers(player1 Player, player2 Player, deck Deck) (Player, Player, Deck) {
	cards1 := deck[0:2]
	cards2 := deck[2:4]
	player1.hand = cards1
	player2.hand = cards2
	return player1, player2, deck[4:]
}

func DoFlop(table Table, deck Deck) (Table, Deck) {
	table.community[0], table.community[1], table.community[2] = deck[0], deck[1], deck[2]
	return table, deck[3:]
}

func DoTurn(table Table, deck Deck) (Table, Deck) {
	table.community[3] = deck[0]
	return table, deck[1:]
}

func DoRiver(table Table, deck Deck) (Table, Deck) {
	table.community[4] = deck[0]
	return table, deck[1:]
}

//func HasRoyalFlush(table Table, player Player) bool {
//
//	return false
//}

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

//func HasStraightFlush() bool {
//	return ""
//}
//
//func HasStraight() bool {
//	return ""
//}
//
//func HasFlush() bool {
//	return ""
//}
//
//func HasFourOfKind() bool {
//	return ""
//}
//
//func HasThreeOfKind() bool {
//	return ""
//}
//
//func HasTwoPair() bool {
//	return ""
//}
//
//func HasPair() bool {
//	return ""
//}
//
