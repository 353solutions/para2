package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	start := time.Now()
	out := PMap(inc, values, 3)
	fmt.Println(out) // {2, 3, 4, 5, 6, 7, 8, 9, 10}
	duration := time.Since(start)

	fmt.Println(duration)

}

// PMap returns application of fn on every value in values.
// It'll run concurrently using up to n goroutines.
// Use golang.org/x/sync/semaphore
func PMap[V any, R any](fn func(V) R, values []V, n int) []R {
	sema := semaphore.NewWeighted(int64(n))
	out := make([]R, len(values))
	ctx := context.TODO()

	for i, v := range values {
		// Go < 1.22
		// i, v := i, v
		sema.Acquire(ctx, 1)
		go func() {
			defer sema.Release(1)
			// Don't need to lock since each goroutine access it's own value in out.
			out[i] = fn(v)
		}()
	}

	sema.Acquire(ctx, int64(n)) // Wait for all goroutines to finish

	return out
}

func inc(n int) int {
	time.Sleep(100 * time.Millisecond)
	return n + 1
}
