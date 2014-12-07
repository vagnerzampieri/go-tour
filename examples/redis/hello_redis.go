package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	connect, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	defer connect.Close()

	connect.Do("SET", "message1", "Hello World")

	world, err := redis.String(connect.Do("GET", "message1"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
}
