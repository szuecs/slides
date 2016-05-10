// show examples: godoc -ex -goroot `pwd` foo
package foo_test

import (
	"fmt"
	"foo"
)

// You can document an example's output, by adding an output comment at its end.
// The output comment must begin with "Output:", as shown below:
func ExampleFoo_HelloA() {
	f := foo.New("Hello A", "B")
	fmt.Println(f.HelloA())
	// Output: Hello A
}

// blah blah
func ExampleFoo_HelloA_2() {
	f := foo.New("A", "B")
	fmt.Println(f.HelloA())
	// Output: A
}
