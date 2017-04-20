package main

import "fmt"

type f func(string) string

func hello(s string) string {
	return "Hello " + s + "!"
}

func main() {
	fmt.Println(hello("Go Class"))
	var myfunc f = hello
	fmt.Println(myfunc("from stored function myfunc"))
	var anononym f = func(s string) string {
		return "Hello from anononym to " + s
	}
	fmt.Println(anononym("Go Class"))

	sum := func(a, b int) int {
		return a + b
	}(2, 3)
	fmt.Printf("sum is: %d\n", sum)
} // END
