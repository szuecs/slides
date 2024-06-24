package main

import "fmt"

var x float64 // package global

func main() {
	var y, z float64 // function local
	x = 3.0 + y
	fmt.Println("x:", x, ", y:", y, ", z:", z)

	const nihongo string = "日本語"

	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

}
