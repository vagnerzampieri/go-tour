package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	csvfile, _ := os.Open("food.csv")

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	rawCSVdata, _ := reader.ReadAll()

	perTypes := make(map[string][]string)
	food := make(map[string][]string)

	for _, each := range rawCSVdata {
		food[each[0]] = append(food[each[0]], each[1])
		perTypes[each[1]] = append(perTypes[each[1]], each[0])
		//fmt.Printf("name: %s, type: %s\n", each[0], each[1])
		//fmt.Printf("perTypes\n", perTypes)
	}

	for k, vs := range food {
		var str string

		str = strings.Join(vs, ", ")

		fmt.Printf("<td alt=\"%s\" style=\"border-bottom: 1px solid #000000; border-left: 1px solid #0000 border-right: none; border-top: 1px solid #000000; cursor: pointer; padding-bottom: 0.04in; padding-left: 0.04in; padding-right: 0in; padding-top: 0.04in;\" title=\"%s\" width=\"14%%\">%s</td>\n", str, str, k)
	}
	fmt.Println(food)
}
