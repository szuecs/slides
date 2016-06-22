package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	s := "abc"
	t := reverse(s)
	fmt.Printf("%s is reversed %s\n", s, t)
}

func reverse(s string) string {
	var buf bytes.Buffer

	for i := len(s) - 1; i >= 0; i-- {
		err := buf.WriteByte(s[i])
		if err != nil {
			log.Fatalf("Can not write '%v' to buffer, caused by: %v", s[i], err)
		}
	}

	return buf.String()
}
