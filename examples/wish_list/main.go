package main

import "fmt"

type WishList []string

func (list WishList) Categorize() ([]string, []string, []string) {
	vegetables, meat, other := []string{}, []string{}, []string{}
	for _, e := range list {
		switch e {
		case "Alface", "Pepino":
			vegetables = append(vegetables, e)
		case "Atum", "Frango":
			meat = append(meat, e)
		default:
			other = append(other, e)
		}
	}
	return vegetables, meat, other
}

func main() {
	list := make(WishList, 6)
	list[0] = "Alface"
	list[1] = "Pepino"
	list[2] = "Azeite"
	list[3] = "Atum"
	list[4] = "Frango"
	list[5] = "Chocolate"

	fmt.Println(list)

	vegetables, meat, other := list.Categorize()
	fmt.Println("Vegetables:", vegetables)
	fmt.Println("Meat:", meat)
	fmt.Println("Other:", other)
}
