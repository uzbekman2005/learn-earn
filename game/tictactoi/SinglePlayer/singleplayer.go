package singleplayer

import (
	"fmt"
	hardpart "learn/game/tictactoi/SinglePlayer/HardPart"
	tictacconsole "learn/game/tictactoi/tictacConsole"
	"learn/globalfunctions"
)

func MainSinglePlayerTicTac() {

	var choose int
	for {
		board := [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
		tictacconsole.SinglePlayerMenu()
		choose = globalfunctions.InputNum(">>>> ")
		if choose == 1 {
			fmt.Println("Easy")
		} else if choose == 2 {
			globalfunctions.SystemClear()
			hardpart.HardMain(board)
		} else {
			break
		}
	}
}
