package main

/*
Concurrency:
- spin work
- spin work, wait for finish
	sync.WaitGroup
- spin work, check error
	- errgroup
- spin work, get result
	- channel
*/

/* Channel semantics
- send/receive blocks until opposite operation(*)
	- guarantee of delivery
	- buffered channel of capacity "n", has "n" non-blocking sends
		- buffer size should be 1
- receive from closed channel returns zero value without blocking
	- use "v, ok :- <- ch" to see if ch is closed
- closing a closed or nil channel will panic
	- channel ownership
- send to a closed channel will panic
- send/receive to/from a nil channel block forever
*/

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
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

	urls := []string{
		"https://go.dev",
		"https://google.com",
		"https://ibm.com/no/such/page",
	}

	// Only wait
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		// wg.Add(1)
		go func() {
			defer wg.Done()
			checkURL(url)
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("%d URLS in %v\n", len(urls), duration)

	// Get result
	start = time.Now()

	ch1 := make(chan result)
	for _, url := range urls {
		go func() {
			r := result{url: url}
			r.status, r.err = urlInfo(url)
			ch1 <- r
		}()
	}

	for range urls {
		r := <-ch1
		fmt.Println(r)
	}

	duration = time.Since(start)
	fmt.Printf("%d URLS in %v\n", len(urls), duration)

	var ch3 chan int

	go func() {
		v := <-ch3
		fmt.Println("v:", v)
	}()

	time.Sleep(100 * time.Millisecond)
	ch3 = make(chan int, 1)
	ch3 <- 3
	time.Sleep(100 * time.Millisecond)

	urlCh := make(chan string)
	var urlWg sync.WaitGroup
	urlWg.Add(len(urls))
	for i := range 2 {
		go poolWorker(i, urlCh, &urlWg)
	}

	for _, url := range urls {
		urlCh <- url
	}
	close(urlCh) // signal no more work

	urlWg.Wait()

	ch4 := make(chan int, 3)
	ch4 <- 1
	ch4 <- 2
	close(ch4)
	// ch4 <- 3 // panic
	v, ok := <-ch4
	fmt.Println(v, ok)
	v, ok = <-ch4
	fmt.Println(v, ok)
	v, ok = <-ch4
	fmt.Println(v, ok)

}

func poolWorker(id int, ch <-chan string, wg *sync.WaitGroup) {
	for url := range ch {
		slog.Info("worker", "id", id, url, url)
		checkURL(url)
		wg.Done()
	}
}

type result struct {
	// call context
	url string

	status string
	err    error
}

func urlInfo(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	return resp.Status, nil
}

func checkURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("checkURL", "url", url, "error", err)
	}

	defer resp.Body.Close()
	slog.Info("checkURL", "url", url, "status", resp.Status)

}
