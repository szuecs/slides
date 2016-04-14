// Package foo is an example that shows the idea of a package.
package foo

// Foo is written capitalized, so it's exported
type Foo struct {
	A string // capitalized member is exported automatically
	b string // not exported
}

func New(a, b string) *Foo { // just a convenient name, but not necessary ;)
	return &Foo{a, b}
}

// HelloA is an exported method
func (f *Foo) HelloA() string {
	return f.A
}

func (f *Foo) helloB() string {
	return f.b
}
