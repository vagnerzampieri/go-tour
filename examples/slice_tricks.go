package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"e"}

	fmt.Println(a)

	a = append(a, "d")
	fmt.Println(a)

	a = append(a, b[0])
	fmt.Println(a)

	fmt.Println(a[1:3])

	fmt.Println(a[2])

	c := []string{}
	d := 2
	e := a[:d]
	f := a[d+1:]
	fmt.Println(e)
	fmt.Println(f)

	c = append(c, e...)
	c = append(c, f...)

	fmt.Println(c)
}
