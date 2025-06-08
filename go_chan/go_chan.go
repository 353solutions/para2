package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("goroutine:", i)
		}()
	}

	fmt.Println("main end")
	time.Sleep(10 * time.Millisecond)

	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	val := <-ch // receive
	fmt.Println("val:", val)

	go func() {
		for i := range 4 {
			msg := fmt.Sprintf("msg #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got:", msg)
	}

	/* The above for loop does:
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("got:", msg)
	}
	*/

	msg := <-ch // ch is closed
	fmt.Printf("msg: %q\n", msg)
	msg, ok := <-ch // ch is closed
	fmt.Printf("msg: %q (ok=%v)\n", msg, ok)
}

/* Channel semantics
- send/receive blocks until opposite operation(*)
	- guarantee of delivery
- receive from closed channel returns zero value without blocking
	- use "v, ok :- <- ch" to see if ch is closed
*/
