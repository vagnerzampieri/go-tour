package main

import (
	"fmt"
	"os"
)

func CreateFiles(dirBase string, files ...string) {
	for _, name := range files {
		filepath := fmt.Sprintf(
			"%s/%s.%s", dirBase, name, "txt")

		file, err := os.Create(filepath)

		defer file.Close()

		if err != nil {
			fmt.Printf("Erro ao criar file %s: %v\n",
				name, err)
			os.Exit(1)
		}

		fmt.Printf("Arquivo %s criado. \n", file.Name())
	}
}

func main() {
	tmp := os.TempDir()

	CreateFiles(tmp)
	CreateFiles(tmp, "test1")
	CreateFiles(tmp, "test2", "test3", "test4")
}
