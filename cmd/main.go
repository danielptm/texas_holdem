package main

import (
	"go_holdem/game"
)

func main() {
	t := game.BuildTable()
	game.PlayHand(t)
}

//x := game.GetChips()
//fmt.Println(x)

//	reader := bufio.NewReader(os.Stdin)
//	reader.ReadString('\n')
//}
