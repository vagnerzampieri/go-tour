package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Connect struct {
	Conn redis.Conn
}

var (
	connect Connect
)

func init() {
	var err error
	connect.Conn, err = redis.Dial("tcp", ":6379")

	if err != nil {
		panic(err)
	}
}

func main() {
	defer connect.Conn.Close()

	connect.Conn.Do("SET", "message1", "Hello World")

	world, err := redis.String(connect.Conn.Do("GET", "message1"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
}
