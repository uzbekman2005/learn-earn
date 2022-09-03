package main

import (
	"errors"
	frontend "learn/frontEnd"
	"learn/game"
	"learn/globalfunctions"
	"learn/registration"
)

var choose int

func main() {
	isRegistrationSucceded := registration.RegistrationMain()
	if isRegistrationSucceded != 1 {
		err := errors.New("registration failed")
		globalfunctions.WriteErrorsToFile(err)
		panic("Registration Failed")
	}
	globalfunctions.SystemClear()
	for {
		frontend.MainMune()
		choose = globalfunctions.InputNum(">>> ")
		if choose == 1 {
			globalfunctions.SystemClear()
			frontend.UserProfile()
		} else if choose == 2 {
			globalfunctions.SystemClear()
			globalfunctions.ShowHighSocers()
		} else if choose == 3 {
			globalfunctions.SystemClear()
			globalfunctions.UpdateUserInfo()
		} else if choose == 4 {
			globalfunctions.SystemClear()
			game.GamesMain() 
		} else {
			break
		}
	}
}
