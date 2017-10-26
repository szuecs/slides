package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Printf("%s", s)
}

func main() {
	go say("hello")
	time.Sleep(1 * time.Second)
	go say(" world!\n")
	time.Sleep(100 * time.Millisecond)
}
