package main

import "fmt"

// T is an example struct
type T struct {
	a, b int
}

func main() {
	var t0 *T = &T{1, 2} // ptr to T
	t0.b = 3
	t1 := &T{3, 4} // type taken from expr
	t1.a = 2
	fmt.Println("t1.a:", t1.a, ", t0.b:", t0.b)
}
