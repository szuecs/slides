// show examples: godoc -ex -goroot `pwd` x
package x_test

import (
	"fmt"
	"x"
)

// You can document an example's output, by adding an output comment at its end.
// The output comment must begin with "Output:", as shown below:
func ExampleFoo_HelloA() {
	f := x.NewFoo("Hello A", "B")
	fmt.Println(f.HelloA())
	// Output: Hello A
}

// blah blah
func ExampleFoo_HelloA_2() {
	f := x.New(Foo"A", "B")
	fmt.Println(f.HelloA())
	// Output: A
}
