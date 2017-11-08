package main

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	s := "abc"
	expected := "cba"
	if res := reverse(s); res != expected {
		t.Errorf("Failed to reverse string, got '%s', expected '%s'", res, expected)
	}
}
