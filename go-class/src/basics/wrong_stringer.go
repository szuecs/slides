package main

import "fmt"

type Foo struct {
	A string
}

// To satisfy the Stringer interface you have to define the receiver
// on the type itself (f Foo) and not on a pointer type (f *Foo).  A
// pointer to type Foo will automatically call String() defined on
// (f Foo), too.
func (f *Foo) String() string {
	return fmt.Sprintf("<Foo.A=%q>", f.A)
}

func main() {
	fo := Foo{"another Foo"}
	fmt.Println(fo) // => {another Foo}
	foo := &fo
	fmt.Println(foo) // => <Foo.A="another Foo">
}
