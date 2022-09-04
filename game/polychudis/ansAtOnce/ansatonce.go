package ansatonce

import (
	graphs "learn/game/polychudis/Graphs"
	"strings"
)

func AnsAtOnceMain(question, answer string) int {
	userAns := graphs.InputString("Enter your answer: ")
	if strings.EqualFold(strings.ToLower(userAns), strings.ToLower(answer)) {
		return 1
	}
	return 0
}
