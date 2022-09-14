package polychudis

import (
	"learn/config"
	"learn/database"
	graphs "learn/game/polychudis/Graphs"
	ansatonce "learn/game/polychudis/ansAtOnce"
	ansbyletter "learn/game/polychudis/ansbyletter"
	question "learn/game/polychudis/questions"
	"learn/globalfunctions"
)

func MainPolichudis() {
	var choice, res, doStop int
	for {
		graphs.MainMenu()
		choice = graphs.InputNum("> ")
		if choice == 1 {
			break
		} else if choice == 2 {
			graphs.SystemClear()
			question1, answer := question.GetRandomQuestion()
			graphs.ShowQuestion(question1)
			graphs.ShowOptions()
			choice = graphs.InputNum(">")
			graphs.SystemClear()
			switch choice {
			case 1:
				res = ansbyletter.AnswerByLetterMain(question1, answer)
			case 2:
				res = ansatonce.AnsAtOnceMain(question1, answer)
			default:
				continue
			}
			doStop = graphs.ShowRes(res)
			if res == 1 {
				config.CurrentUser.Score += 10
				config.CurrentUser.HighScore += 10
				database.SetInfo(config.CurrentUser, config.Client)
			}
			if doStop == 1 {
				globalfunctions.SystemClear()
				continue
			} else {
				break
			}

		} else {
			break
		}

	}
}
