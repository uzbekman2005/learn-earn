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
			return LogInMain()
		} else {
			break
		}
	}
	return -1
}

func SignUpMain() int {
	config.CurrentUser.First_name = globalfunctions.InputString("First name: ")
	config.CurrentUser.Last_name = globalfunctions.InputString("Last name: ")
	config.CurrentUser.Username = GetUserName()
	config.CurrentUser.Password = globalfunctions.InputString("Password: ")
	config.CurrentUser.Score = 0
	config.CurrentUser.HighScore = 0
	config.CurrentUser.Country = globalfunctions.InputString("Country: ")
	config.CurrentUser.DateOfBirth = globalfunctions.InputString("Date of birth: ")
	globalfunctions.WriteNewUserToFile()
	globalfunctions.WriteToUserstxt()
	frontend.SuccesSignUp()

	return 1
}

func GetUserName() string {
	tempName := ""
	for {
		frontend.UsernameRequirement()
		tempName = globalfunctions.InputString("Username: ")
		if globalfunctions.IsValidUserName(tempName) {
			return tempName
		}
	}
}

func LogInMain() int {
	for {
		frontend.LogInMenu()
		userNameinput := globalfunctions.InputString("Username: ")
		password := globalfunctions.InputString("Password: ")
		if globalfunctions.IsRightLogin(userNameinput, password) {
			err := globalfunctions.ReadSpeicificUser(userNameinput)
			globalfunctions.CheckErr(err)
			if err == nil {
				frontend.SuccesLogin()
				return 1
			} else {
				return -1
			}
		}
		frontend.WasRegistered()
		choose = globalfunctions.InputNum("> ")
		if choose == 2{
			frontend.SignUpMenu()
			return SignUpMain()
		}

	}
}
