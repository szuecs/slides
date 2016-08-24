package main

import "fmt"

type Foo struct {
	A string
}

// To make it work use (f Foo) to not define it on pointers.  A
// pointer to type Foo will automatically call String() defined on
// (f Foo), too.
func (f *Foo) String() string {
	return fmt.Sprintf("<Foo.A=%q>", f.A)
}

func main() {
	fo := Foo{"another Foo"}
	fmt.Println(fo) // => <Foo.A="another Foo">
}
