package registration

import (
	"errors"
	"learn/config"
	"learn/database"
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
			if err := LogInMainWithDb(&config.CurrentUser); err == nil {
				globalfunctions.WriteErrorsToFile(err)
				return 1
			} else {
				globalfunctions.SystemClear()
				frontend.WasRegistered()
			}
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
	//Writing to database
	err := database.SetInfo(config.CurrentUser, config.Client)
	globalfunctions.CheckErr(err)
	//Writing to File
	frontend.SuccesSignUp()
	return 1
}

func LogInMainWithDb(user *config.User) error {
	globalfunctions.SystemClear()
	frontend.LogInMenu()
	userNameinput := globalfunctions.InputString("Username: ")
	password := globalfunctions.InputString("Password: ")
	if database.IsInDataBase(userNameinput, password, config.Client) {
		err := database.GetInfo(user, userNameinput, config.Client)
		globalfunctions.CheckErr(err)
		return nil
	} else {
		globalfunctions.SystemClear()
		return errors.New("INVALID DATA")
	}
}

/* This Function is not working anymore
func LogInMain(isSecon bool) int {
	globalfunctions.SystemClear()
	for {
		frontend.LogInMenu()
		userNameinput := globalfunctions.InputString("Username: ")
		password := globalfunctions.InputString("Password: ")

		if globalfunctions.IsRightLogin(userNameinput, password) && userNameinput != config.CurrentUser.Username {

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
*/
