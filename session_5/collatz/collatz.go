package collatz

import (
	"math/rand/v2"
	"sync"
	"time"
)

func Collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}

	return n*3 + 1
}

func CollatzLen(n int) int {
	// Add some jitter
	time.Sleep(time.Duration(rand.IntN(100)) * time.Microsecond)

	count := 1
	for n != 1 {
		count++
		n = Collatz(n)
	}

	return count
}

func MaxCollatz(n int) (int, int) {
	var mu sync.Mutex
	maxN, maxLen := 0, 0

	sema := make(chan bool, 10)
	var wg sync.WaitGroup

	wg.Add(n - 1)
	for i := 1; i < n; i++ {
		sema <- true
		go func() {
			defer func() {
				<-sema
				wg.Done()
			}()

			s := CollatzLen(i)
			mu.Lock()
			defer mu.Unlock()
			if s > maxLen {
				maxN, maxLen = i, s
			}
		}()
	}

	wg.Wait()
	return maxN, maxLen
}
