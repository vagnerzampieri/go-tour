package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

	for typ, names := range perTypes {
		fmt.Println("<table><tbody>")
		fmt.Println(typ, "<br />")
		for i, n := range names {
			str := strings.Join(food[n], ", ")

			if i == 0 {
				fmt.Println("<tr>")
				fmt.Printf("<td alt=\"%s\" style=\"border: 1px solid #000000; cursor: pointer; padding: 0.04in;\" title=\"%s\" width=\"14%%\">%s</td>\n", str, str, n)
			} else if (i+1) == len(names) || (i+1)%3 == 0 {
				fmt.Printf("<td alt=\"%s\" style=\"border: 1px solid #000000; cursor: pointer; padding: 0.04in;\" title=\"%s\" width=\"14%%\">%s</td>\n", str, str, n)
				fmt.Println("</tr>")
				fmt.Println("<tr>")
			} else {
				fmt.Printf("<td alt=\"%s\" style=\"border: 1px solid #000000; cursor: pointer; padding: 0.04in;\" title=\"%s\" width=\"14%%\">%s</td>\n", str, str, n)
			}
			//fmt.Println(i, str)
		}
		fmt.Println("</table></tbody>")
	}

	//fmt.Println(food)
}
