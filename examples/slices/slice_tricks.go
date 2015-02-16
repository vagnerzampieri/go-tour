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

	u := make([]int, 10, 20)
	fmt.Println(u, len(u), cap(u))

	original := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", original)

	newSlice := original[1:3]
	fmt.Println("Novo:", newSlice)

	original[2] = 13

	fmt.Println("Original pos modificacao:", original)
	fmt.Println("Novo pos modificacao:", newSlice)

	fmt.Println("Capacidade original:", cap(original))
	fmt.Println("Capacidade modificado:", cap(newSlice))

	fmt.Println("Tamanho original:", len(original))
	fmt.Println("Tamanho modificado:", len(newSlice))
}
