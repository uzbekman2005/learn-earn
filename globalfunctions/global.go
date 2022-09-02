package globalfunctions

import (
	"encoding/json"
	"fmt"
	"learn/config"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// var err error
// var file os.File

func WriteErrorsToFile(err error) {
	file, errr := os.OpenFile("/home/azizbek/go/src/Projects/learn-earn/Users/AllUsers/errors.txt", os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errr != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%v, Error type: %v\n", time.Now(), err))
}

func InputNum(hint string) int {
	var temp string

	for {
		fmt.Print(hint)
		fmt.Scan(&temp)
		res, err := strconv.Atoi(temp)
		if err == nil {
			return res
		}
		fmt.Println("Enter only INT num!!!")
	}
}

func InputString(hint string) string {
	fmt.Print(hint)
	var temp string
	fmt.Scan(&temp)
	return temp
}

func CheckErr(err error) {
	if err != nil {
		WriteErrorsToFile(err)
		fmt.Println(err)
	}
}

func UpdateUserInfo() {
	// This Function will be writtend afte person structis created
	// func update Menu is Ready
}

func IsValidUserName(username string) bool {
	if !unicode.IsLetter(rune(username[0])) {
		return false
	}
	for i := 0; i < len(username); i++ {
		if !unicode.IsDigit(rune(username[i])) && !unicode.IsLetter(rune(username[i])) && username[i] != '_' {
			return false
		}
	}
	return true && IsUniqueUserName(username)
}

func IsUniqueUserName(username string) bool {
	filename := "/home/azizbek/go/src/Projects/learn-earn/Users/AllUsers/users.txt"
	info, err := os.ReadFile(filename)
	CheckErr(err)
	temp := strings.Split(string(info), "\n")
	for _, elem := range temp {
		oneUser := strings.Split(string(elem), ",")
		if oneUser[0] == username {
			return false
		}
	}
	return true
}

func WriteNewUserToFile() {
	filename := fmt.Sprintf("/home/azizbek/go/src/Projects/learn-earn/Users/Individualuser/%s.json", config.CurrentUser.Username)
	data, err := json.Marshal(config.CurrentUser)
	CheckErr(err)
	err = os.WriteFile(filename, data, 0644)
	CheckErr(err)
}

// This Function Writes username and password to the users.txt when new user signed up
func WriteToUserstxt() {
	filename := "/home/azizbek/go/src/Projects/learn-earn/Users/AllUsers/users.txt"
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		panic(err)
	}
	CheckErr(err)
	f.WriteString(fmt.Sprintf("%s,%s\n", config.CurrentUser.Username, config.CurrentUser.Password))
}

func IsRightLogin(username, password string) bool {
	filename := "/home/azizbek/go/src/Projects/learn-earn/Users/AllUsers/users.txt"
	info, err := os.ReadFile(filename)
	CheckErr(err)
	temp := strings.Split(string(info), "\n")
	for _, elem := range temp {
		oneUser := strings.Split(string(elem), ",")
		if oneUser[0] == username && password == oneUser[1]{
			return true
		}
	}
	return false
}

// This function read from Users/Infividualusers/username.json and 
// converts it to struct and initialize it to currentUser
func ReadSpeicificUser(username string) error {
	filename := fmt.Sprintf("/home/azizbek/go/src/Projects/learn-earn/Users/Individualuser/%s.json", username)
	content, err := os.ReadFile(filename)
	CheckErr(err)
	err = json.Unmarshal(content,  &config.CurrentUser)
	return err
}