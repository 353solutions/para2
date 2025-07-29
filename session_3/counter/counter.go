package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// go run -race counter.go
// -race is also supported by build and test
// Common practice: Test with -race
func main() {
	const nGr = 10
	var wg sync.WaitGroup
	wg.Add(nGr)

	/* Solution 1: Use mutex
	var mu sync.Mutex
	counter := 0
	*/
	// Solution 2: sync/atomic
	var counter atomic.Int64

	for range nGr {
		go func() {
			defer wg.Done()
			for range 1000 {
				time.Sleep(time.Microsecond)
				/* Solution 1
				mu.Lock()
				counter++ // Read, modify, write (not atomic)
				mu.Unlock()
				*/
				counter.Add(1)
			}
		}()
	}

	wg.Wait()
	// fmt.Println("counter:", counter)
	fmt.Println("counter:", counter.Load())
}
