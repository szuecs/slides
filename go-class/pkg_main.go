// .
// └── src
//     └── foo
//         └── a.go
// ├── pkg_main.go
// $ GOPATH=`pwd` go run pkg_main.go
// A
// $ cat pkg_main.go
package main

import (
	"foo"
)

func main() {
	f := foo.New("A", "b")
	println(f.HelloA())
	// Compile time error: f.helloB()
	println(f.helloB()) // f.helloB undefined (cannot refer to unexported field or method
	// foo.(*Foo)."".helloB)
}
