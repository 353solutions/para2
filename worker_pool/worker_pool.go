package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	tasks := make(chan Task[string])

	for range 1 {
		go worker(tasks)
	}

	t := Task[string]{
		fn: uuid,
		ch: make(chan string, 1),
	}
	go func() {
		for range 4 {
			tasks <- t
		}
	}()

	// work ...
	// when need result
	v := <-t.ch
	fmt.Println("v:", v)

	time.Sleep(100 * time.Millisecond)

}

func Future[T any](fn func() T) chan T {
	ch := make(chan T, 1)

	go func() {
		v := fn()
		ch <- v
	}()

	return ch
}

func uuid() string {
	// example of work
	buf := make([]byte, 16)
	rand.Read(buf)
	return fmt.Sprintf("%x", buf)
}

func worker[T any](ch chan Task[T]) {
	for t := range ch {
		fmt.Println("worker start")
		v := t.fn()
		t.ch <- v
		fmt.Println("worker end")
	}
}

type Task[T any] struct {
	fn func() T
	ch chan T
}
