package main

import "fmt"

func main() {
	sliceInteger := []int{1, 2, 3, 5, 7, 11, 13}
	fmt.Println("sliceInteger ==", sliceInteger)

	for i := 0; i < len(sliceInteger); i++ {
		fmt.Printf(
			"sliceInteger[%d] == %d\n",
			i,
			sliceInteger[i])
	}

	sliceString := []string{"Anaheim Ducks", "Detroit Red Wings", "Pittsburgh Penguins"}
	fmt.Printf("sliceString == %s\n", sliceString)

	for i := 0; i < len(sliceString); i++ {
		fmt.Printf(
			"sliceString[%d] == %s\n",
			i,
			sliceString[i])
	}

	sliceByte := []byte{'a', 1}
	fmt.Printf("sliceByte == %v\n", sliceByte)
}
