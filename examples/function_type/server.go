package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc(
		"/time",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s",
				time.Now().UTC().Format("2006-01-02 15:04:05"))
		})

	http.ListenAndServe(":8080", nil)
}
