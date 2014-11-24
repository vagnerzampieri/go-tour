package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
	fmt.Println("m[\"Bell Labs\"] -> ", m["Bell Labs"])
	fmt.Println("m[\"Bell Labs\"].Lat ->", m["Bell Labs"].Lat)
	fmt.Println("m[\"Bell Labs\"].Long ->", m["Bell Labs"].Long)
	fmt.Println("m[\"Google\"] ->", m["Google"])
	fmt.Println("m[\"Google\"].Lat ->", m["Google"].Lat)
	fmt.Println("m[\"Google\"].Long ->", m["Google"].Long)
}
