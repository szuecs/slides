package main

import "fmt"

func set(i int, dst *int) {
	*dst = i
}

func swap(i, j *int) {
	*i, *j = *j, *i
}

func main() {
	var i, j int
	i = 0
	j = 1
	fmt.Printf("i: %d, j: %d\n", i, j)
	set(42, &i)
	fmt.Printf("i: %d, j: %d\n", i, j)
	swap(&i, &j)
	fmt.Printf("i: %d, j: %d\n", i, j)
}
