package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, String("I'm a frayed knot."))
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, &Struct{"Hello", ":", "Gophers!"})
}

func main() {
	var s Struct
	var str String

	http.Handle("/string", str)
	http.Handle("/struct", s)
	http.ListenAndServe(":4000", nil)
}
