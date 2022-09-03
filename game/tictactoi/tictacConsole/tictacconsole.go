package tictacconsole

import "fmt"

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
	fmt.Println("|   1. Login second user                             |")
	fmt.Println("|   2. Start the game                                |")
	fmt.Println("|   3. Exit                                          |")
	fmt.Println("+----------------------------------------------------+")
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
