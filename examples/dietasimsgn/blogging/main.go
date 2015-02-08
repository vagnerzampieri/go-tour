package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var ws []string

	fmt.Println("Título do arquivo: ")
	filename, _ := reader.ReadString('\n')
	filename = strings.Split(filename, "\n")[0]

	fmt.Println("Dificuldade: ")
	difficulty, _ := reader.ReadString('\n')
	ws = append(ws, difficulty)

	fmt.Println(filename)
	fmt.Println(difficulty)
	/*
		fmt.Println("Tempo de preparo: ")
		preparationTime, _ := reader.ReadString('\n')
		fmt.Println(preparationTime)

		fmt.Println("Tempo de cozimento: ")
		cookingTime, _ := reader.ReadString('\n')
		fmt.Println(cookingTime)

		fmt.Println("Rendimento: ")
		foodYield, _ := reader.ReadString('\n')
		fmt.Println(foodYield)

		fmt.Println("Ingredientes: ")
		ingredients, _ := reader.ReadString('\n')
		fmt.Println(ingredients)

		fmt.Println("Modo de preparo: ")
		methodOfPreparation, _ := reader.ReadString('\n')
		fmt.Println(methodOfPreparation)

		fmt.Println("Exemplo para combinações: ")
		exampleCombinations, _ := reader.ReadString('\n')
		fmt.Println(exampleCombinations)

		fmt.Println("Pode ser usado no preparo?: ")
		canBeUsedInCooking, _ := reader.ReadString('\n')
		fmt.Println(canBeUsedInCooking)
	*/
	output, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	for w := range ws {
		if _, err := output.WriteString(ws[w]); err != nil {
			panic(err)
		}
	}

}
