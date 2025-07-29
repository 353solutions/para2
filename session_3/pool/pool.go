package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// chanPoolDemo()
	pool := sync.Pool{
		// New: NewConn, // Will create new one when there are no free (unlimited connections)
	}
	for i := range 3 {
		c := Conn{i}
		runtime.AddCleanup(&c, func(_ bool) { c.Close() }, true)
		pool.Put(&c)
	}

	for i := range 10 {
		go poolWorker(i, &pool)
	}

	time.Sleep(time.Second)
}

func poolWorker(id int, pool *sync.Pool) {
	for {
		v := pool.Get()
		if v == nil {
			fmt.Println("no conn available")
			time.Sleep(100 * time.Millisecond)
			continue
		}

		conn := v.(*Conn)
		fmt.Printf("%d: using %d\n", id, conn.ID)
		time.Sleep(100 * time.Millisecond)
		pool.Put(conn)
	}
}

var connID int64

func NewConn() any {
	i := atomic.AddInt64(&connID, 1)
	c := Conn{int(i)}
	return &c
}

func chanPollDemo() {
	pool := NewChanPool(3)
	for i := range 10 {
		go chanPoolWorker(i, pool)
	}

	time.Sleep(time.Second)
}

func (c ChanPool) Release(conn Conn) {
	c.ch <- conn
}

func (c ChanPool) Acquire() Conn {
	return <-c.ch
}

func NewChanPool(size int) *ChanPool {
	ch := make(chan Conn, size)
	for i := range size {
		ch <- Conn{i}
	}

	p := ChanPool{
		ch: ch,
	}
	return &p
}

type ChanPool struct {
	ch chan Conn
}

// TOOD: Context
func chanPoolWorker(id int, pool *ChanPool) {
	for {
		conn := pool.Acquire()
		fmt.Printf("%d: using %d\n", id, conn.ID)
		time.Sleep(100 * time.Millisecond)
		pool.Release(conn)
	}
}

func (c *Conn) Close() error {
	return nil
}

type Conn struct {
	ID int
}
