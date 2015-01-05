package main

import "fmt"

type File struct {
	name      string
	length    float64
	caracters int
	words     int
	lines     int
}

func main() {
	file := File{"artigo.txt", 12.68, 12986, 1862, 220}
	fmt.Println(file)
	fmt.Printf("%s\t%.2fKB\n", file.name, file.length)

	stackFile := &File{length: 7.29, name: "file.txt"}
	fmt.Printf("before: %s\t%.2fKB\n", stackFile.name, stackFile.length)

	stackFile.length = 17.47
	fmt.Printf("after: %s\t%.2fKB\n", stackFile.name, stackFile.length)
}
