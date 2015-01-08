package main

import (
	"fmt"
	"io"
)

type ReaderStrings struct{}

func (r ReaderStrings) Read(p []byte) (int, error) {
	p[0] = 'A'
	p[1] = 'B'
	p[2] = 'C'
	p[3] = 'D'

	return len(p), nil
}

func readString(r io.Reader) string {
	p := make([]byte, 4)
	r.Read(p)
	return string(p)
}

func main() {
	reader := ReaderStrings{}
	fmt.Println(reader)
	fmt.Println(readString(reader))
}
