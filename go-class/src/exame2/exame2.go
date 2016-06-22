// call it from a main
package main

import "fmt"
import "rstring"

func main() {
	rs := rstring.New("hello world")
	fmt.Println(rs.Reverse())
}
