package redisexample

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func GetStruct(c redis.Conn) error {

	const objectPrefix string = "user:"

	username := "otto"
	s, err := redis.String(c.Do("GET", objectPrefix+username))
	if err == redis.ErrNil {
		fmt.Println("User does not exist")
	} else if err != nil {
		return err
	}

	usr := User{}
	err = json.Unmarshal([]byte(s), &usr)

	fmt.Printf("%+v\n", usr)

	return nil

}
