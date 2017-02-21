package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.Buffer{}
	buf.Write([]byte("Hello"))
	buf.Write([]byte(" "))
	buf.Write([]byte("World!"))
	s := buf.String()
	fmt.Println(s)
}
