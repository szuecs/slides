// .
// └── src
//     └── x
//         └── a.go
// ├── pkg_main.go
// $ GOPATH=`pwd` go run pkg_main.go
// A
// $ cat pkg_main.go
package main

import (
	"x"
)

func main() {
	f := x.NewFoo("A", "b")
	println(f.HelloA())
	// Compile time error: f.helloB()
	println(f.helloB()) // f.helloB undefined (cannot refer to unexported field or method
	// foo.(*Foo)."".helloB)
}
