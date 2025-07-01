package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	t := NewToken()
	defer t.Close()

	for i := range 4 {
		// i := i // Go < 1.22
		go func() {
			for range time.Tick(300 * time.Millisecond) {
				fmt.Printf("%d: %q\n", i, t.Value())
			}
		}()
	}

	time.Sleep(4 * time.Second)
}

func (t *Token) Close() error {
	t.cancel()

	// Close return error to implement io.Closer
	return nil
}

func NewToken() *Token {
	// Design question: Do we want to pass ctx to NewToken?
	ctx, cancel := context.WithCancel(context.Background())
	t := Token{
		value:  refreshToken(),
		cancel: cancel,
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				tok := refreshToken()

				t.mu.Lock()
				t.value = tok
				t.mu.Unlock()
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()

	return &t
}

/* Exercise:
You should refresh the token every second in another goroutine
Use sync.RWMutex

Harder:
Have only single goroutine that reads,writes the data. Communicate with it via channels
*/

func (t *Token) Value() string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.value
}

type Token struct {
	cancel context.CancelFunc
	mu     sync.RWMutex
	value  string
}

// 3rd party package, you can't change it
var tokN = 0

func refreshToken() string {
	tokN++
	fmt.Println("refresh")
	return fmt.Sprintf("token %d", tokN)
}
