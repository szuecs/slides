package rstring

import (
	"bytes"
	"log"
)

type RString struct {
	S string
}

// START
func New(s string) RString {
	return RString{S: s}
}

func (s RString) Reverse() RString {
	var buf bytes.Buffer

	for i := len(s.S) - 1; i >= 0; i-- {
		err := buf.WriteByte(s.S[i])
		if err != nil {
			log.Fatalf("Can not write '%v' to buffer, caused by: %v", s.S[i], err)
		}
	}

	return New(buf.String())
} // END
