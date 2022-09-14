package hardpart

import "math"

const (
	player   = "X"
	computer = "O"
)

func IsMovesLeft(board [][]string) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == "-" {
				return true
			}
		}
	}
	return false
}

func Evaluate(board [][]string) int {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][2] == board[i][0] {
			if board[i][0] == player {
				return -10
			} else {
				return 10
			}
		}
	}

	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[2][i] == board[0][i] {
			if board[0][i] == player {
				return -10
			} else {
				return 10
			}
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[0][0] == player {
			return -10
		} else {
			return 10
		}
	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		if board[0][2] == player {
			return -10
		} else {
			return 10
		}
	}

	return 0
}

func MiniMax(board [][]string, depth int, isMax bool) int {
	score := Evaluate(board)

	if score == 10 || score == -10 {
		return score
	}

	if !IsMovesLeft(board) {
		return 0
	}

	m := len(board)
	n := len(board[0])

	if isMax {
		best := -1000
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] == "-" {
					board[i][j] = computer
					best = int(math.Max(float64(best), float64(MiniMax(board, depth+1, !isMax))))
					board[i][j] = "-"
				}
			}
		}
		return best
	} else {
		best := 1000
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] == "-" {
					board[i][j] = computer
					best = int(math.Min(float64(best), float64(MiniMax(board, depth+1, !isMax))))
					board[i][j] = "-"
				}
			}
		}
	}
	return 0
}

func FindBestMove(board [][]string) (resi int, resj int) {
	bestVal := -1000
	resi = -1
	resj = -1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				board[i][j] = computer
				moveVal := MiniMax(board, 0, false)
				board[i][j] = "-"
				if moveVal > bestVal {
					bestVal = moveVal
					resi = i
					resj = j
				}
			}
		}
	}
	return
}
