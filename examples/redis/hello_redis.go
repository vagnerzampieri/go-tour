package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Connect struct {
	c redis.Conn
}

var (
	connect Connect
)

func init() {
	var err error
	connect.c, err = redis.Dial("tcp", ":6379")

	if err != nil {
		panic(err)
	}
}

func main() {
	defer connect.c.Close()

	connect.c.Do("SET", "message1", "Hello World")

	world, err := redis.String(connect.c.Do("GET", "message1"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
}
