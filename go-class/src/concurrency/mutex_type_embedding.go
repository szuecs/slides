package main

import (
	"fmt"
	"sync"
	"time"
)

//START
type Counter struct {
	sync.Mutex // type embedding
	n          int
}

func (c *Counter) inc() {
	c.Lock()
	c.n++
	c.Unlock()
}

func main() {
	syncedCounter := &Counter{}
	for i := 1; i < 10; i++ {
		go func() { // goroutine + anonymous function
			syncedCounter.inc()
			fmt.Printf("n: %d\n", syncedCounter.n)
		}() // invoke
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("n: %d\n", syncedCounter.n)
}

//END
