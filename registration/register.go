package registration

import (
	"learn/config"
	frontend "learn/frontEnd"
	"learn/globalfunctions"
)

var choose int

func RegistrationMain() int {
	for {
		frontend.RegistrationMenu()
		choose = globalfunctions.InputNum("Enter you Choice: ")
		if choose == 1 {
			return SignUpMain()
		} else if choose == 2 {
			return LogInMain(false)
		} else {
			break
		}
	}
	return -1
}

func SignUpMain() int {
	config.CurrentUser.First_name = globalfunctions.InputString("First name: ")
	config.CurrentUser.Last_name = globalfunctions.InputString("Last name: ")
	config.CurrentUser.Username = globalfunctions.GetUserName()
	config.CurrentUser.Password = globalfunctions.GetPassword()
	config.CurrentUser.Score = 0
	config.CurrentUser.HighScore = 0
	config.CurrentUser.Country = globalfunctions.InputString("Country: ")
	config.CurrentUser.DateOfBirth = globalfunctions.InputDate("Date of birth: ")
	globalfunctions.WriteNewUserToFile()
	globalfunctions.WriteToUserstxt()
	frontend.SuccesSignUp()
	return 1
}

func LogInMain(isSecon bool) int {
	globalfunctions.SystemClear()
	for {
		frontend.LogInMenu()
		userNameinput := globalfunctions.InputString("Username: ")
		password := globalfunctions.InputString("Password: ")
		if globalfunctions.IsRightLogin(userNameinput, password) {
			err := globalfunctions.ReadSpeicificUser(userNameinput, isSecon)
			globalfunctions.CheckErr(err)
			if err == nil {
				frontend.SuccesLogin()
				return 1
			} else {
				return -1
			}
		}
		globalfunctions.SystemClear()
		frontend.WasRegistered()
		choose = globalfunctions.InputNum("> ")
		if choose == 2 {
			frontend.SignUpMenu()
			return SignUpMain()
		}

	}
}
