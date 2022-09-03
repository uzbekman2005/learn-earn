package game

import (
	frontend "learn/frontEnd"
	polychudis "learn/game/poly-chudis"
	"learn/globalfunctions"
)

func GamesMain() {
	var choose int

	for {
		globalfunctions.SystemClear()
		frontend.GamesMenu()
		choose = globalfunctions.InputNum("> ")
		if choose == 1{
			polychudis.MainPolichudis()
		}else{
			break
		}
	}
}