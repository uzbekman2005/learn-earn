package main

import (
	"errors"
	"fmt"
	frontend "learn/frontEnd"
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

	for {
		frontend.MainMune()
		choose = globalfunctions.InputNum(">>> ")
		if choose == 1 {
			frontend.UserProfile()
		} else if choose == 2 {
			fmt.Println("HighScore will be displayed")
		} else if choose == 3 {
			globalfunctions.UpdateUserInfo()
		} else if choose == 4 {
			fmt.Println("Play Game")
		} else {
			break
		}
	}
}
