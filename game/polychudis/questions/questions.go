package questions

import (
	"learn/globalfunctions"
	"math/rand"
	"os"
	"strings"
)

func GetRandomQuestion() (question, answer string) {
	questions := getAllQuestions()
	randVal := rand.Intn(len(questions)-1) + len(question) - 1
	temp := strings.Split(questions[randVal], "|")
	question = temp[0]
	answer = temp[1]
	return
}

func getAllQuestions() []string {
	data, err := os.ReadFile("/home/azizbek/go/src/Projects/learn-earn/game/polychudis/questions/q_and_ans.txt")
	globalfunctions.CheckErr(err)
	return strings.Split(string(data), "\n")
}
