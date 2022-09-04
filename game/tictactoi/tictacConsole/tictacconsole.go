package tictacconsole

import (
	"fmt"
)
var IsSecondUserRegistered bool

func TicMainMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("|           Welcome to Tic Tac Toi game              |")
	fmt.Println("|    We can offer you different levels of Tic game   |")
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("|                 Choose what you want               |")
	fmt.Println("|     1. Single Player                               |")
	fmt.Println("|     2. Multi Player                                |")
	fmt.Println("|     3. Exit                                        |")
	fmt.Println("+----------------------------------------------------+")
}

func MultiPlayerMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("|       In this part second user should log in       |")
	fmt.Println("+----------------------------------------------------+")
	if IsSecondUserRegistered {
		fmt.Printf("|   1. Login second user %-4s                       |\n", "✅")
	} else {
		fmt.Println("|   1. Login second user                             |")
	}
	fmt.Println("|   2. Start the game                                |")
	fmt.Println("|   3. Exit                                          |")
	fmt.Println("+----------------------------------------------------+")
}

func MultiPlayerShowPlayerInfo(name1, name2 string, score1, score2 int, turn1, turn2 string) {
	fmt.Println("+------------------+-------+ +-------+------------------+")
	fmt.Printf("|  %-15s | %-5d | | %-5d |  %-15s |\n", name1, score1, score2, name2)
	fmt.Printf("|     X            | %-5s| | %-5s|    O             |\n", turn1, turn2)
	fmt.Println("+------------------+-------+ +-------+------------------+")
}

func ShowBoard(board [][]string) {
	fmt.Println("+-----+-----+-----+")
	for _, row := range board {
		fmt.Printf("|  %-3s|  %-3s|  %-3s|\n", string(row[0]), string(row[1]), string(row[2]))
		fmt.Println("+-----+-----+-----+")
	}
}

func SinglePlayerMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("|        Here are the levels of the game             |")
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("|    1. Easy                                         |")
	fmt.Println("|    2. Hard                                         |")
	fmt.Println("|    3. Exit                                         |")
	fmt.Println("+----------------------------------------------------+")
}
