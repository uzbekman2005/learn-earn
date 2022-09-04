package ansbyletter

import (
	"fmt"
	graphs "learn/game/polychudis/Graphs"
	"strings"
)

var outputAnswer []string

func AnswerByLetterMain(question, answer string) int {
	outputAnswer = generateOutPutAnswer(len(answer), answer)
	rightAnswer := strings.Split(answer, "")
	numberOfRounds := int(float64(len(rightAnswer)) * 1.5)
	graphs.ShowQuestion(question)
	for numberOfRounds > 0 {
		showAnswer(outputAnswer)
		ans := graphs.InputString("Enter a letter")
		isRightAnswer(ans, rightAnswer)
		graphs.SystemClear() //console is cleared here
		graphs.ShowQuestion(question)
		if strings.Join(outputAnswer, "") == answer {
			return 1
		}
		numberOfRounds--
	}
	return 0
}

func isRightAnswer(letter string, rightAnswer []string) {
	for i := 0; i < len(rightAnswer); i++ {
		temp := strings.ToLower(rightAnswer[i])
		if temp == strings.ToLower(letter) && outputAnswer[i] == "-" {
			outputAnswer[i] = rightAnswer[i]
		}
	}
}

func generateOutPutAnswer(l int, rightAnswer string) []string {
	var res []string
	for i := 0; i < l; i++ {
		if string(rightAnswer[i]) != " " {
			res = append(res, "-")
		} else {
			res = append(res, " ")
		}
	}
	return res
}

func showAnswer(list []string) {
	var temp []string
	for i := 0; i < len(list); i++ {
		temp = append(temp, "---")
	}
	fmt.Print("+", strings.Join(temp, "+"), "+\n")
	fmt.Println("|", strings.Join(list, " | "), "|")
	fmt.Print("+", strings.Join(temp, "+"), "+\n")

}
