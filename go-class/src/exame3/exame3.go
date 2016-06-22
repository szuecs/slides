// call it from a main
package main

import (
	"fmt"
	"sort"
)

import "rstring"

func main() {
	rs := rstring.New("hello world")
	fmt.Println(rs.Reverse())
	rs1 := rstring.New("a")
	rs2 := rstring.New("ab")
	rs3 := rstring.New("abc")

	ary := rstring.SortableRString{rs, rs1, rs3, rs2}
	fmt.Printf("%v\n", ary)
	sort.Sort(ary)
	fmt.Printf("%v\n", ary)

}
