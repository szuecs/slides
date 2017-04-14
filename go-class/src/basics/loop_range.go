package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5} // define a slice

	// index loop (use only condition to get a while loop)
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	// range loop
	for idx, item := range a {
		fmt.Printf("At index %d, there is item %d\n", idx, item)
	}
}
