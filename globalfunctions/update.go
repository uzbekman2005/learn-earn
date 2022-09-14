package globalfunctions

import (
	"learn/config"
	"learn/database"
	frontend "learn/frontEnd"
	"strconv"
	"strings"
)

var choose int

// Using this funtion user can update his/her info
func UpdateUserInfo() {
	for {
		SystemClear()
		frontend.UpdateMenu()
		choose = InputNum(">>> ")
		if choose == 1 {
			config.CurrentUser.First_name = InputString("First name: ")
		} else if choose == 2 {
			config.CurrentUser.Last_name = InputString("Last name: ")
		} else if choose == 3 {
			config.CurrentUser.Country = InputString("Country: ")
		} else if choose == 4 {
			config.CurrentUser.DateOfBirth = InputDate("Date of birth ")
		} else if choose == 5 {
			updatePassword()
		} else if choose == 6 {
			username := config.CurrentUser.Username
			UpdateUsername(config.CurrentUser.Password)
			if err := config.Client().Del(username).Err(); err != nil {
				CheckErr(err)
			}
		} else {
			// Changes are written to files user.txt username.json
			SystemClear()
			break
		}
		// WriteNewUserToFile(config.CurrentUser.Username, false)
		if err := database.SetInfo(config.CurrentUser, config.Client); err != nil {
			CheckErr(err)
		}
	}
}
func GetUserName() string {
	tempName := ""
	for {
		frontend.UsernameRequirement()
		tempName = InputString("Username: ")
		if IsValidUserName(tempName) {
			return tempName
		}
	}
}
func GetPassword() string {
	temp := ""
	for {
		frontend.PasswordRequirement()
		temp = InputString("Password: ")
		if ValidPassword(temp) {
			return temp
		}
	}
}

func ValidPassword(password string) bool {
	// Password to be valid it should be at least 5 in length
	// don't contain commas
	if len(password) < 5 || strings.Contains(password, ",") || strings.Contains(password, " ") {
		return false
	}
	return true

}

func UpdateUsername(password string) {
	for {
		newUsername := GetUserName()
		if !database.IsInDataBase(newUsername, password, config.Client) {
			config.CurrentUser.Username = newUsername
			return
		} else {
			SystemClear()
		}
	}
}

// This Function is used to update password
func updatePassword() {
	newPassword := GetPassword()
	config.CurrentUser.Password = newPassword
}

// This function is used to input date info correctly
func InputDate(hint string) string {
	temp := ""
	for {
		frontend.DateFormat()
		temp = InputString(hint)
		if IsValidDate(temp) {
			return temp
		}
	}
}

// This function checks whether the date is valid or not
func IsValidDate(date string) bool {
	arrDate := strings.Split(date, "-")
	if len(arrDate) != 3 {
		return false
	}
	year, err := strconv.Atoi(arrDate[0])
	if err != nil {
		CheckErr(err)
		return false
	}
	month, err := strconv.Atoi(arrDate[1])
	CheckErr(err)
	day, err := strconv.Atoi(arrDate[2])
	CheckErr(err)
	if year > 9999 || year < 0 || month > 12 || month <= 0 || day > 31 || day < 1 {
		return false
	}

	if !isLeapYear(year) && month == 2 && day == 29 {
		return false
	}
	if month == 2 && day > 29 {
		return false
	}
	if month == 4 || month == 6 || month == 9 || month == 11 {
		if day > 30 {
			return false
		}
	}
	return true
}

// This Function check whether the year is leap year
func isLeapYear(year int) bool {
	return (year%100 != 0 && year%4 == 0) || year%400 == 0
}
