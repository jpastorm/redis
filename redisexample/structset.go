package redisexample

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

type User struct {
	Username  string
	MobileID  int
	Email     string
	FirstName string
	LastName  string
}

func SetStruct(c redis.Conn) error {

	const objectPrefix string = "user:"

	usr := User{
		Username:  "otto",
		MobileID:  1234567890,
		Email:     "ottoM@repoman.com",
		FirstName: "Otto",
		LastName:  "Maddox",
	}

	// serialize User object to JSON
	json, err := json.Marshal(usr)
	if err != nil {
		return err
	}

	// SET object
	_, err = c.Do("SET", objectPrefix+usr.Username, json)
	if err != nil {
		return err
	}

	return nil
}
