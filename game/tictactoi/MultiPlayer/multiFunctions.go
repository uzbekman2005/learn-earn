package multiplayer

import (
	"learn/config"
	"learn/database"
	tictacconsole "learn/game/tictactoi/tictacConsole"
	"learn/globalfunctions"
	"learn/registration"
)

// This function is used to log in or sign up second user
func LoginSeconduser() int {
	if err := registration.LogInMainWithDb(&config.SecondUser); err != nil {
		return 0
	}
	return 1
}

var board [][]string

func MultiPlayerPlay() int {
	board = [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
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
		tictacconsole.ShowBoard(board)
		answer = globalfunctions.InputNum(">>> ")
		if rounds%2 == 1 {
			isNext = PutTheAnswer(board, answer, "X")
		} else {
			isNext = PutTheAnswer(board, answer, "O")
		}
		if !isNext {
			continue
		}
		res := IsWin(board)
		if res == 1 {
			globalfunctions.SystemClear()
			if rounds%2 == 0 {
				config.SecondUser.Score += 20
				config.SecondUser.HighScore += 20
				tictacconsole.YouWin(config.SecondUser.First_name)
				database.SetInfo(config.SecondUser, config.Client)
			} else {
				config.CurrentUser.Score += 20
				config.CurrentUser.HighScore += 20
				tictacconsole.YouWin(config.CurrentUser.First_name)
				database.SetInfo(config.CurrentUser, config.Client)
			}
			return 1
		} else if IsEqual(board) {
			return 1
		}
		rounds++

	}
}

func IsWin(board [][]string) int {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][2] == board[i][0] && board[i][0] != "-" {
			return 1
		}
		if board[2][i] == board[1][i] && board[2][i] == board[0][i] && board[0][i] != "-" {
			return 1
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != "-" {
		return 1
	}
	if board[1][1] == board[0][2] && board[0][2] == board[2][0] && board[1][1] != "-" {
		return 1
	}
	return 0
}

func PutTheAnswer(board [][]string, num int, ans string) bool {
	if num < 1 || num > 9 {
		return false
	}
	num--
	x := num / 3
	y := num % 3
	if board[x][y] == "-" {
		board[x][y] = ans
		return true
	}
	return false
}

func IsEqual(board [][]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	config.CurrentUser.Score += 10
	config.CurrentUser.HighScore += 10

	config.SecondUser.Score += 10
	config.SecondUser.HighScore += 10
	if err := database.SetInfo(config.CurrentUser, config.Client); err != nil {
		globalfunctions.CheckErr(err)
	}
	err := database.SetInfo(config.SecondUser, config.Client)
	globalfunctions.CheckErr(err)
	return true
}
