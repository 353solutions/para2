package main

import (
	"fmt"
	"time"
)

func main() {
	t := NewToken()
	for i := range 4 {
		go func() {
			for range time.Tick(300 * time.Millisecond) {
				fmt.Printf("%d: %s\n", i, t.Value())
			}
		}()
	}

	time.Sleep(4 * time.Second)
}

/* Exercise:
You should refresh the token every second in another goroutine
Hint: sync.RWMutex
*/

func (t *Token) Value() string {
	return "TODO"
}

type Token struct {
	value string
}

var tokN = 0

func refreshToken() string {
	tokN++
	fmt.Println("refresh")
	return fmt.Sprintf("token %d", tokN)
}
