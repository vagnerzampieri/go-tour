package main

import "fmt"

func main() {
	var a [3]int
	numbers := [5]int{1, 2, 3, 4, 5}
	prime := [...]int{2, 3, 5, 7, 11, 13}
	names := [2]string{}

	fmt.Println(a, numbers, prime, names)
	fmt.Println(len(a), len(numbers), len(prime), len(names))

	// Valor inicial varia de acordo com o tipo de dado definido para o array
	// false para valores do tipo bool
	// 0 para ints
	// 0.0 para floats
	// ""(ou string vazia) para strings
	// nil para ponteiros, funcoes, interfaces, slices, maps e channels

	var multiA [2][2]int
	fmt.Println("Multi:", multiA)

	multiA[0][0], multiA[0][1] = 3, 5
	multiA[1][0], multiA[1][1] = 7, -2

	multiB := [2][2]int{{2, 13}, {-1, 6}}

	fmt.Println("Multi A:", multiA)
	fmt.Println("Multi B:", multiB)
}
