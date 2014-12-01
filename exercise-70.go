package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	fmt.Println("ch1 ->", ch1)
	fmt.Println("ch2 ->", ch2)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	// tree.New(2)
	ch := make(chan int)
	fmt.Println("ch ->", ch)
	fmt.Println("tree.New(1) ->", tree.New(1))
	fmt.Println("tree.New(2) ->", tree.New(2))

	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Println("v ->", v)
	}

	fmt.Println("Same(tree.New(1), tree.New(1)) ->", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2)) ->", Same(tree.New(1), tree.New(2)))
}
