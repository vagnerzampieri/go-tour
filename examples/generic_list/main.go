package main

import "fmt"

type GenericList []interface{}

func (list *GenericList) RemoveIndex(index int) interface{} {
	l := *list
	remove := l[index]
	*list = append(l[0:index], l[index+1:]...)
	return remove
}

func (list *GenericList) RemoveFirst() interface{} {
	return list.RemoveIndex(0)
}

func main() {
	list := GenericList{
		"Anaheim Ducks", 42, true, 9, "New Orleans Saints", 3.14,
	}

	fmt.Printf("Lista:\n%v\n\n", list)
}
