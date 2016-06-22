package rstring

import (
	"bytes"
	"log"
)

type RString struct {
	S string
}

func New(s string) *RString {
	return &RString{S: s}
}

func (s *RString) Reverse() *RString {
	var buf bytes.Buffer

	for i := len(s.S) - 1; i >= 0; i-- {
		err := buf.WriteByte(s.S[i])
		if err != nil {
			log.Fatalf("Can not write '%v' to buffer, caused by: %v", s.S[i], err)
		}
	}

	return New(buf.String())
}

// START
func (s *RString) String() string {
	return s.S
}

type SortableRString []*RString

// implement the sort interface
func (s SortableRString) Len() int {
	return len(s)
}
func (s SortableRString) Less(i, j int) bool {
	return s[i].S < s[j].S
}
func (s SortableRString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//END
