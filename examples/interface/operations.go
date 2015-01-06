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

func (a Age) String() string {
	return fmt.Sprintf("Idade desde %d", a.birthdate)
}

func main() {

}
