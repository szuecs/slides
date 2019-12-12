package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.Buffer{}
	buf.Write([]byte("Hello"))
	buf.WriteByte(32)
	buf.WriteString("World!")
	s := buf.String()
	fmt.Println(s)
}
