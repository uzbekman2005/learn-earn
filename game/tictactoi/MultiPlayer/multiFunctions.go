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
		res := IsWin(Board)
		if res == 1 {
			globalfunctions.SystemClear()
			if rounds%2 == 0 {
				config.SecondUser.Score += 20
				config.SecondUser.HighScore += 20
				tictacconsole.YouWin(config.SecondUser.First_name)
				globalfunctions.WriteNewUserToFile(config.SecondUser.Username, true)
			} else {
				config.CurrentUser.Score += 20
				config.CurrentUser.HighScore += 20
				tictacconsole.YouWin(config.CurrentUser.First_name)
				globalfunctions.WriteNewUserToFile(config.CurrentUser.Username, false)
			}
			return 1
		}else if IsEqual(Board){
			return 1
		}
		rounds++
		if rounds == 10 {
			break
		}
	}
	return 1
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
	globalfunctions.WriteNewUserToFile(config.CurrentUser.Username, false)
	globalfunctions.WriteNewUserToFile(config.SecondUser.Username, true)
	return true
}
