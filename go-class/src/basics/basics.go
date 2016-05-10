package main

import "fmt"

var x float64 // package global

func main() {
	var y, z float64 // function local
	x = 3.0 + y
	fmt.Println("x:", x, ", z:", z)

	const str string = "this is a 日本語 string\n"
	var ch byte = str[11]
	r := []rune(str)
	fmt.Printf("ch: %q, rune: %q\n", ch, r[11])
}
