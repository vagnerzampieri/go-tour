package main

import "fmt"

type GenericList []interface{}

func (list *GenericList) RemoveIndex(index int) interface{} {
	l := *list
	remove := l[index]
	*list = append(l[0:index], l[index+1:]...)
	return remove
}
