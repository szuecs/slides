package main

import "fmt"

func main() {
	var x, y float64
	x = 3.0 + y
	fmt.Println("x:", x)

	type T struct{ a, b int }
	var t0 *T = &T{1, 2} // ptr to T
	t0.b = 3
	t1 := &T{3, 4} // type taken from expr
	t1.a = 2
	fmt.Println("t1.a:", t1.a, ", t0.b:", t0.b)

	const str string = "this is a 日本語 string\n"
	if len(str) > 0 {
		var ch byte = str[11]
		r := []rune(str)
		fmt.Printf("ch: %q, rune: %q\n", ch, r[11])
	}
}
