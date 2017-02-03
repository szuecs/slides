package main

import "fmt"

type Foo struct {
	A string
}

// implements stringer Interface
func (f Foo) String() string {
	return fmt.Sprintf("<Foo.A=%q>", f.A)
}

func main() {
	fo := Foo{"another Foo"}
	fmt.Println(fo) // => <Foo.A="another Foo">
	foo := &fo
	fmt.Println(foo) // => <Foo.A="another Foo">
}
