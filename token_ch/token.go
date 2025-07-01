package main

import (
	"context"
	"fmt"
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

type Cmd struct {
	op    string
	value string
	ch    chan string
}

func NewToken() *Token {
	// Design question: Do we want to pass ctx to NewToken?
	ctx, cancel := context.WithCancel(context.Background())
	t := Token{
		ch:     make(chan Cmd),
		cancel: cancel,
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				tok := refreshToken()
				t.ch <- Cmd{"write", tok, nil}
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()

	tok := refreshToken()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case cmd := <-t.ch:
				if cmd.op == "read" {
					cmd.ch <- tok
					continue
				}
				tok = cmd.value
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
	cmd := Cmd{
		op: "read",
		// buffered so reader won't block
		ch: make(chan string, 1),
	}

	t.ch <- cmd
	return <-cmd.ch
}

type Token struct {
	ch     chan Cmd
	cancel context.CancelFunc
}

// 3rd party package, you can't change it
var tokN = 0

func refreshToken() string {
	tokN++
	fmt.Println("refresh")
	return fmt.Sprintf("token %d", tokN)
}
