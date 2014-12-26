package main

import "fmt"

type WishList []string

func printSlice(slice []string) {
	fmt.Println("Slice:", slice)
}

func printList(list WishList) {
	fmt.Println("Lista de compras:", list)
}

func main() {
	list := WishList{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"}

	printSlice([]string(list))
	printList(WishList(slice))
}
