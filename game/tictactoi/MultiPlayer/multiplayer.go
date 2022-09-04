package multiplayer

import (
	"learn/config"
	tictacconsole "learn/game/tictactoi/tictacConsole"
	"learn/globalfunctions"
)

func MultiplayerMain() {
	globalfunctions.SystemClear()
	choose := -1
	tictacconsole.IsSecondUserRegistered = false
	for {
		tictacconsole.MultiPlayerMenu()
		choose = globalfunctions.InputNum("> ")
		if choose == 1 {
			if LoginSeconduser() == 1 {
				tictacconsole.IsSecondUserRegistered = true
			}
		} else if choose == 2 {
			if !tictacconsole.IsSecondUserRegistered || config.SecondUser.Username == config.CurrentUser.Username {
				globalfunctions.SystemClear()
				continue
			}
			res := MultiPlayerPlay()
			if res == 1 {
				continue
			}
		} else {
			break
		}
	}

}
