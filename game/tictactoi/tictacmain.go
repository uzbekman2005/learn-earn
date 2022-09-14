package tictactoi

import (
	multiplayer "learn/game/tictactoi/MultiPlayer"
	singleplayer "learn/game/tictactoi/SinglePlayer"
	tictacconsole "learn/game/tictactoi/tictacConsole"
	"learn/globalfunctions"
)

func TicTacMain() {
	globalfunctions.SystemClear()
	choose := 1
	for {
		tictacconsole.TicMainMenu()
		choose = globalfunctions.InputNum("> ")
		if choose == 1 {
			singleplayer.MainSinglePlayerTicTac()
		} else if choose == 2 {
			multiplayer.MultiplayerMain()
		} else {
			break
		}
	}
}
