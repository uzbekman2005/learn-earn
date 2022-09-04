package game

import (
	frontend "learn/frontEnd"
	polychudis "learn/game/polychudis"
	"learn/game/tictactoi"
	"learn/globalfunctions"
)

func GamesMain() {
	var choose int

	for {
		globalfunctions.SystemClear()
		frontend.GamesMenu()
		choose = globalfunctions.InputNum("> ")
		if choose == 1 {
			polychudis.MainPolichudis()
		} else if choose == 2 {
			tictactoi.TicTacMain()
		} else {
			break
		}
	}
}
