package config

import "github.com/go-redis/redis"

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

func Client() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
