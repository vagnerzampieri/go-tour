package main

import "fmt"

type State struct {
	name       string
	population int
	capital    string
}

func main() {
	states := make(map[string]State, 2)
	states["GO"] = State{"Goias", 6434052, "Goiania"}
	states["PB"] = State{"Paraiba", 3914418, "Joao Pessoa"}

	fmt.Println("States ->", states)
	fmt.Println("GO ->", states["GO"])
	fmt.Println("GO name ->", states["GO"].name)
	fmt.Println("GO population ->", states["GO"].population)
	fmt.Println("GO capital ->", states["GO"].capital)

	fmt.Println("SP ->", states["SP"])

	sp, found := states["SP"]
	if found {
		fmt.Println("SP found", sp)
	}
}
