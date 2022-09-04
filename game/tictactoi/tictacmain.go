package tictactoi

import (
	"fmt"
	multiplayer "learn/game/tictactoi/MultiPlayer"
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
			fmt.Println("Single Player")
		} else if choose == 2 {
			multiplayer.MultiplayerMain()
		} else {
			break
		}
	}
}
