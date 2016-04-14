package main

import (
	"fmt"
	"sync"
)

var syncedCounter struct {
	sync.Mutex // type embedding
	n          int
}

func inc() {
	syncedCounter.Lock()
	syncedCounter.n++
	syncedCounter.Unlock()
}

func main() {
	fmt.Printf("n: %d\n", syncedCounter.n)

	for i := 1; i < 10; i++ {
		go func() { // goroutine + anonymous function
			inc()
			fmt.Printf("n: %d\n", syncedCounter.n)
		}() // invoke
	}

	fmt.Printf("n: %d\n", syncedCounter.n)
}
