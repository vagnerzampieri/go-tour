package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The vale", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The vale", m["Answer"])

	fmt.Println("m ->", m)

	delete(m, "Answer")
	fmt.Println("The value", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	if !ok {
		m["Answer"] = 40
		fmt.Println("The value:", m["Answer"])
	}
}
