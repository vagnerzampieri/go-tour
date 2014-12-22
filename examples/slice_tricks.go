package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"e"}

	fmt.Println("a ->", a)

	a = append(a, "d")
	fmt.Println("append(a, 'd') ->", a)

	a = append(a, b[0])
	fmt.Println("append(a, b[0]) ->", a)

	fmt.Println("a[1:3] ->", a[1:3])

	fmt.Println("a[2] ->", a[2])

	c := []string{}
	d := 2
	e := a[:d]
	f := a[d+1:]
	fmt.Println("a[:d] ->", e)
	fmt.Println("a[d+1:] ->", f)

	c = append(c, e...)
	c = append(c, f...)

	fmt.Println("append(c, e...) append(c, f...) ->", c)

	var s []int
	names := []string{}
	fmt.Println(s, names)

	t := make([]int, 10)
	fmt.Println(t, len(t), cap(t))
}
