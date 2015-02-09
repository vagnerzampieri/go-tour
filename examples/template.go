package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}} \n")
	if err != nil {
		panic(err)
	}
	out, _ := os.Create("baba.txt")
	defer out.Close()
	err = tmpl.Execute(out, sweaters)
	if err != nil {
		panic(err)
	}
}
