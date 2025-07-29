package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for i := range 3 {
			ch1 <- fmt.Sprintf("[1]: %d", i)
			time.Sleep(time.Microsecond)
		}
		close(ch1)
	}()
	go func() {
		for i := range 4 {
			ch2 <- fmt.Sprintf("[2]: %d", i)
			time.Sleep(time.Microsecond)
		}
		close(ch2)
	}()

	for v := range FanIn([]chan string{ch1, ch2}) {
		fmt.Println(v)
	}

}

func FanIn(chans []chan string) chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	wg.Add(len(chans))
	go func() {
		for _, ch := range chans {
			go func() {
				defer wg.Done()
				for v := range ch {
					out <- v

				}
			}()
		}
		wg.Wait()
		close(out)
	}()

	return out
}
