package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"learn/config"

	"github.com/go-redis/redis"
)

func SetInfo(user config.User, client func() *redis.Client) error {
	jsonInfo, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	clnt := client()
	err = clnt.Set(user.Username, string(jsonInfo), 0).Err()
	return err
}

func GetInfo(user *config.User, username string, client func() *redis.Client) error {
	clnt := client()
	usr := clnt.Get(username).Val()
	err := errors.New("ERORR WHILE GETTING DATA FROM DATABASE")
	
	if user == nil {
		return err
	}
	err = json.Unmarshal([]byte(usr), &user)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func IsInDataBase(username, password string, client func() *redis.Client) bool {
	clnt := client()
	res := clnt.Exists(username).Val()
	if res == 0 {
		return false
	}
	var temp config.User

	err := GetInfo(&temp, username, client)
	if err != nil {
		fmt.Println(err)
	}
	return temp.Password == password
}

func IsUniqueUserName(username string, client func() *redis.Client) bool {
	clnt := client()
	if res := clnt.Exists(username).Val(); res == 1 {
		return false
	}
	return true
}
