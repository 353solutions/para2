package main

import (
	"fmt"
	"time"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	start := time.Now()
	out := PMap(inc, values, 3)
	duration := time.Since(start)

	fmt.Println(out, duration)

}

// PMap returns application of fn on every value in values concurrently.
// It'll use up to n goroutines
// Use golang.org/x/sync/semaphore
func PMap[V any, R any](fn func(V) R, values []V, n int) []R {
	return nil

}

func inc(n int) int {
	time.Sleep(100 * time.Millisecond)
	return n + 1
}
