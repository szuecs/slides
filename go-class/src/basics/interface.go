package main

import "fmt"

type Foo interface {
	Say(s string) string
}

type T string

func (t *T) Say(s string) string {
	return fmt.Sprintf("%s world with base %s", s, *t)
}

func foo(f Foo) {
	var s string = f.Say("Hello")
	fmt.Printf("%v", s)
}

func main() {
	var t T = "x"
	foo(&t)
}
