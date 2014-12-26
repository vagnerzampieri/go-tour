package main

import "fmt"

type WishList []string

func main() {
	list := make(WishList, 6)
	list[0] = "Alface"
	list[1] = "Pepino"
	list[2] = "Azeite"
	list[3] = "Atum"
	list[4] = "Frango"
	list[5] = "Chocolate"

	fmt.Println(list)
}
