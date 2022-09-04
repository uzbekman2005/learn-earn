package config

type User struct {
	First_name  string
	Last_name   string
	Username    string
	Password    string
	Score       int
	HighScore   int
	Country     string
	DateOfBirth string
}

var CurrentUser User
var SecondUser User
