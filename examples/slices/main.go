package main

import (
	"fmt"
	"strings"
)

type LinesOfText [][]string

func main() {
	var text LinesOfText
	str := "isso  tem um; ponto e vírgula;outro ; e  outro"
	for _, s := range strings.Split(str, ";") {
		var strs []string
		for _, st := range strings.Split(s, " ") {
			if st != "" {
				strs = append(strs, st)
			}
		}
		var tex = LinesOfText{[]string{strings.Join(strs, " ")}}
		text = append(text, tex...)
	}

	fmt.Println(text)
	fmt.Println(len(text))

	for _, s := range text {
		fmt.Println(s[0])
	}
	/*
		text := LinesOfText{
			[]string{"Now is the time"},
			[]string{"for all good gophers"},
			[]string{"to bring some fun to the party."},
		}

		var tex = LinesOfText{[]string{"End"}}
		fmt.Println(tex)
		text = append(text, tex...)
		fmt.Println(text)
	*/
}
