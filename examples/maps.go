package main

import "fmt"

func main() {
	empty1 := map[int]string{}
	empty2 := make(map[int]string)

	fmt.Println("empty1 ->", empty1)
	fmt.Println("empty2 ->", empty2)

	bigMap := make(map[int]string, 4096)

	fmt.Println("len bigMap ->", len(bigMap))

	capitals := map[string]string{
		"RJ": "Rio de Janeiro",
		"SP": "Sao Paulo",
		"PR": "Curitiba"}

	fmt.Println("Len capitals ->", len(capitals))
	fmt.Println("Capitals ->", capitals)
	fmt.Println("RJ ->", capitals["RJ"])

}
