// Package foo is an example that shows the idea of a package.
package foo

// Foo is written capitalized, so it's exported
type Foo struct {
	A string // capitalized member is exported automatically
	b string // not exported
}

// New is an exported package function that returns the reference of a newly created object Foo
func New(a, b string) *Foo {
	return &Foo{a, b}
}

// HelloA is an exported method
func (f *Foo) HelloA() string {
	return f.A
}

func (f *Foo) helloB() string {
	return f.b
}
