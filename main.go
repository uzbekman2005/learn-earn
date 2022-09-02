package main

import (
	"errors"
	"fmt"
	"learn/config"
	"learn/globalfunctions"
	"learn/registration"
)

func main() {
	isRegistrationSucceded := registration.RegistrationMain()
	if isRegistrationSucceded != 1 {
		err := errors.New("registration failed")
		globalfunctions.WriteErrorsToFile(err)
		panic("Registration Failed")
	}
	fmt.Println("Succesfully Signed in")
	fmt.Println(config.CurrentUser)
}
