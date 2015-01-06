package main

import (
	"fmt"
	"time"
)

type Age struct {
	birthdate int
}

func (a Age) Calc() int {
	return time.Now().Year() - a.birthdate
}

func main() {

}
