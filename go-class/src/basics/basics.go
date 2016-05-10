package main

import "fmt"

var x float64 // package global

func main() {
	var y, z float64 // function local
	x = 3.0 + y
	fmt.Println("x:", x, ", y:", y, ", z:", z)

	const str string = "this is a 日本語 string\n"
	var ch byte = str[12]
	r := []rune(str)
	fmt.Printf("ch: %q, rune: %q\n", ch, r[12])
	fmt.Printf("ch  : %[1]q/%[1]v/%016[1]b\nrune: %[2]q/%[2]v/%016[2]b\n", ch, r[12])
}
