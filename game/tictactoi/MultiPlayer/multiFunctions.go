package multiplayer

import (
	"learn/config"
	tictacconsole "learn/game/tictactoi/tictacConsole"
	"learn/globalfunctions"
	"learn/registration"
)

// This function is used to log in or sign up second user
func LoginSeconduser() int {
	return registration.LogInMain(true)
}

var Board [][]string

func MultiPlayerPlay() int {
	Board = [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	// ans := "0"
	name1 := config.CurrentUser.First_name
	name2 := config.SecondUser.First_name
	score1 := config.CurrentUser.Score
	score2 := config.SecondUser.Score
	rounds := 1
	answer := 0
	isNext := false
	for {
		globalfunctions.SystemClear()
		if rounds%2 == 1 {
			tictacconsole.MultiPlayerShowPlayerInfo(name1, name2, score1, score2, "⏳", "🎧")
		} else {
			tictacconsole.MultiPlayerShowPlayerInfo(name1, name2, score1, score2, "🎧", "⏳")
		}
		tictacconsole.ShowBoard(Board)
		answer = globalfunctions.InputNum(">>> ")
		if rounds%2 == 1 {
			isNext = PutTheAnswer(answer, "X")
		} else {
			isNext = PutTheAnswer(answer, "O")
		}
		if !isNext {
			continue
		}
		rounds++
		if rounds == 10 {
			break
		}
	}
	return 1
}

func PutTheAnswer(num int, ans string) bool {
	if num < 1 || num > 9 {
		return false
	}
	num--
	x := num / 3
	y := num % 3
	if Board[x][y] == "-" {
		Board[x][y] = ans
		return true
	}
	return false
}
