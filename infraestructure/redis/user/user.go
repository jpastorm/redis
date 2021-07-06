package user

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jpastorm/redis/model"
)

// User struct that implement the interface domain.user.Storage
type User struct {
	conn redis.Conn
}

// New return a new User
func New(conn redis.Conn) User {
	return User{conn: conn}
}

// Create this method creates a model.User in postgres db
func (u User) Create(m *model.User) error {
	const objectPrefix string = "user:"
	// SET object
	do, err := u.conn.Do("HMSET", objectPrefix+m.Email, m.Email,m.Password)
	if err != nil {
		return err
	}
	fmt.Println(do)
	return nil
}