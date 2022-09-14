package hardpart

import (
	"learn/config"
	"learn/database"
	multiplayer "learn/game/tictactoi/MultiPlayer"
	tictacconsole "learn/game/tictactoi/tictacConsole"
	"learn/globalfunctions"
)

func HardMain(board [][]string) {
	answer := 1
	rounds := 1
	name := config.CurrentUser.First_name
	isNext := false

	for {
		globalfunctions.SystemClear()
		if rounds%2 == 1 {
			tictacconsole.MultiPlayerShowPlayerInfo(name, "Computer", config.CurrentUser.Score, 0, "🎧", "⏳")
		} else {
			tictacconsole.MultiPlayerShowPlayerInfo(name, "Computer", config.CurrentUser.Score, 0, "⏳", "🎧")
		}
		tictacconsole.ShowBoard(board)
		if rounds%2 == 1 {
			answer = globalfunctions.InputNum(">>> ")
			isNext = multiplayer.PutTheAnswer(board, answer, "X")
		} else {
			answerI, answerJ := FindBestMove(board)
			isNext = multiplayer.PutTheAnswer(board, (answerI*3)+answerJ+1, "O")
		}
		if !isNext {
			continue
		}
		res := multiplayer.IsWin(board)
		if res == 1 {
			if rounds%2 == 1 {
				config.CurrentUser.Score += 50
				config.CurrentUser.HighScore += 50
				database.SetInfo(config.CurrentUser, config.Client)
				tictacconsole.YouWin(config.CurrentUser.First_name)
			} else {
				tictacconsole.YouLost()
			}
			return
		} else if multiplayer.IsEqual(board) {
			tictacconsole.Draw()
			config.CurrentUser.Score += 25
			config.CurrentUser.HighScore += 25
			database.SetInfo(config.CurrentUser, config.Client)
			return
		}
		rounds++
	}

}
