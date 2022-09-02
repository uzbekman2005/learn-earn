package globalfunctions

import (
	"fmt"
	"io/fs"
	"learn/config"
	frontend "learn/frontEnd"
	
	"os"
	"strconv"
	"strings"
)

var choose int

// Using this funtion user can update his/her info
func UpdateUserInfo() {
	for {
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
			updatePassword(config.CurrentUser.Username)
		} else if choose == 6 {
			//Username will be changed using UpdateUsername
			UpdateUsername(config.CurrentUser.Password)
		} else {
			// Changes are written to files user.txt username.json
			WriteNewUserToFile()
			break
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
func UpdateUsername(password string) {
	newUsername := GetUserName()
	filename := "/home/azizbek/go/src/Projects/learn-earn/Users/AllUsers/users.txt"
	content, err := os.ReadFile(filename)
	CheckErr(err)
	contentstr := string(content)
	oldText := fmt.Sprintf("%s,%s", config.CurrentUser.Username, password)
	replaceTxt := fmt.Sprintf("%s,%s", newUsername, password)
	contentstr = strings.ReplaceAll(contentstr, oldText, replaceTxt)
	os.WriteFile(filename, []byte(contentstr), fs.FileMode(os.O_WRONLY))
	// WriteNewUserToFile()
	oldFilename := fmt.Sprintf("/home/azizbek/go/src/Projects/learn-earn/Users/Individualuser/%s.json", config.CurrentUser.Username)
	newFileName := fmt.Sprintf("/home/azizbek/go/src/Projects/learn-earn/Users/Individualuser/%s.json", newUsername)
	err = os.Rename(oldFilename, newFileName)
	CheckErr(err)
	config.CurrentUser.Username = newUsername
	fmt.Println("Success")
}

// This Function is used to update password
func updatePassword(username string) {
	newPassword := InputString("New password: ")
	filename := "/home/azizbek/go/src/Projects/learn-earn/Users/AllUsers/users.txt"
	content, err := os.ReadFile(filename)
	CheckErr(err)
	contentstr := string(content)
	oldText := fmt.Sprintf("%s,%s", username, config.CurrentUser.Password)
	replaceTxt := fmt.Sprintf("%s,%s", username, newPassword)
	contentstr = strings.ReplaceAll(contentstr, oldText, replaceTxt)
	os.WriteFile(filename, []byte(contentstr), fs.FileMode(os.O_WRONLY))
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
