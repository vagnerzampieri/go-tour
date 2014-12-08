package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Repository struct {
	Conn redis.Conn
}

var (
	repo Repository
)

func init() {
	var err error
	repo.Conn, err = redis.Dial("tcp", ":6379")

	if err != nil {
		panic(err)
	}
}

func main() {
	defer repo.Conn.Close()

	repo.Conn.Do("SET", "message1", "Hello World")

	world, err := redis.String(repo.Conn.Do("GET", "message1"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
}
