package main

type Foo struct {
	A string
	b string
}

func (f *Foo) HelloA() string {
	return f.A
}
func (f *Foo) helloB() string {
	return f.b
}

func main() {
	foo := &Foo{"This is A", "This is b"}
	println(foo.HelloA())
	println(foo.helloB())
}
