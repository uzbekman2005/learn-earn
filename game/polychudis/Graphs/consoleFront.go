package graphs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func MainMenu() {
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("+-              Welcome                           -+")
	fmt.Println("+-      I am happy to play with you :)            -+")
	fmt.Println("+--------------------------------------------------+")
	fmt.Print("\n")
	fmt.Println("   Press 1 -> To exit   ")
	fmt.Println("   Press 2 -> To start  ")
	// 1 Start
	// 2 End
}

func ShowOptions() {
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("+-      You have 2 options to answer              -+")
	fmt.Println("+-                                                -+")
	fmt.Println("+-    1. Answer letter by letter                  -+")
	fmt.Println("+-    2. Answer at once                           -+")
	fmt.Println("+--------------------------------------------------+")
}

func ShowRes(res int) int {
	// if 1 win else if 0 lose
	if res == 1 {
		fmt.Println("+---------------------------------------------+")
		fmt.Println("+    Congratulations you won the game!!!     -+")
		fmt.Println("+---------------------------------------------+")
		fmt.Print("\n")
		fmt.Println(" 1. To play again")
		fmt.Println(" 2. To exit the game")
	} else {

		fmt.Println("************************************************")
		fmt.Println("#         Unfortunately you didn't win!!!      #")
		fmt.Println("************************************************")
		fmt.Print("\n")
		fmt.Println(" 1. To play again ")
		fmt.Println(" 2. To exit the game")
	}
	return InputNum(">")
}


func InputNum(hint string) int {
	var temp string

	for {
		fmt.Print(hint)
		fmt.Scan(&temp)
		res, err := strconv.Atoi(temp)
		if err == nil {
			return res
		}
		fmt.Println("Enter only int number")
	}

}

func ShowQuestion(question string) {
	listQuestion := strings.Split(question, " ")
	// total width is 52 bytes
	curLength := 0
	l := len(listQuestion)
	var temp []string
	fmt.Println("+--------------------------------------------------+")
	for i := 0; i < l; i++ {
		if curLength+len(listQuestion[i])+len(temp) > 45 {
			fmt.Printf("+- %-46s -+\n", strings.Join(temp, " "))
			temp = []string{listQuestion[i]}
			curLength = len(listQuestion[i])
		} else {
			curLength += len(listQuestion[i])
			temp = append(temp, listQuestion[i])
		}
	}
	if len(temp) > 0 {
		fmt.Printf("+- %-46s -+\n", strings.Join(temp, " "))
	}
	fmt.Println("+--------------------------------------------------+")
}

func SystemClear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func InputString(hint string) string {
	fmt.Print(hint)
	var temp string
	fmt.Scan(&temp)
	return temp
}
