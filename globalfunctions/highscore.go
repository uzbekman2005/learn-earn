package globalfunctions

import (
	"encoding/json"
	"fmt"
	"learn/config"
	"os"
	"strings"
)

var allUsers []config.User


func ShowHighSocers() {
	ListUsers()
	sortAllUsers()
	// info is shown in table format
	fmt.Println("+---------------+----------------+-------------+----------+")
	fmt.Println("|  Name         |  Country       |  Birth date |  Score   |")
	fmt.Println("+---------------+----------------+-------------+----------+")
	for i := 0; i < len(allUsers) && i < 3; i++{
		fmt.Printf("| %-14s| %-14s | %-11s | %-8d |\n", allUsers[i].First_name, allUsers[i].Country, allUsers[i].DateOfBirth, allUsers[i].Score)
		fmt.Println("+---------------+----------------+-------------+----------+")
	}
}

func sortAllUsers() {
	max := 0
	for i := 0; i < len(allUsers)-1; i++ {
		max = i
		for j := i + 1; j < len(allUsers); j++ {
			if allUsers[i].Score > allUsers[max].Score {
				max = i
			}
		}
		if max != i {
			allUsers[i], allUsers[max] = allUsers[max], allUsers[i]
		}
	}
}

func GetAllUsernames() []string {
	filename := "/home/azizbek/go/src/Projects/learn-earn/Users/AllUsers/users.txt"
	info, err := os.ReadFile(filename)
	CheckErr(err)
	temp := strings.Split(string(info), "\n")
	var res []string
	for i := 0; i < len(temp) -1; i++ {
		elem := temp[i]
		oneUser := strings.Split(string(elem), ",")
		res = append(res, oneUser[0])
	}
	return res
}

func ListUsers() {
	listUsers := GetAllUsernames()
	for i := 0; i < len(listUsers); i++ {
		elem := listUsers[i]
		filename := fmt.Sprintf("/home/azizbek/go/src/Projects/learn-earn/Users/Individualuser/%s.json", string(elem))
		content, err := os.ReadFile(filename)
		CheckErr(err)
		var temp config.User
		err = json.Unmarshal(content, &temp)
		CheckErr(err)
		allUsers = append(allUsers, temp)
	}
}
