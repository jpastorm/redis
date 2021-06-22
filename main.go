package main

import (
	"fmt"

	"github.com/jpastorm/redis/redisexample"
)

func main() {
	pool := redisexample.NewPool()
	conn := pool.Get()
	defer conn.Close()
	err := redisexample.Ping(conn)
	if err != nil {
		fmt.Println(err)
	}
}
