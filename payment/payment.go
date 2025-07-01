package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	p := Payment{
		From:   "Wile E. Coyote",
		To:     "ACME",
		Amount: 1323,
	}
	p.Execute(time.Now())
	p.Execute(time.Now())
}

func (p *Payment) Execute(t time.Time) {
	p.once.Do(func() {
		p.execute(t)
	})
}

func (p *Payment) execute(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("%s: %s -> [%d¢] -> %s\n", ts, p.From, p.Amount, p.To)
}

type Payment struct {
	once sync.Once // see also sync.OnceFunc, sync.OnceValue, sync.OnceValues

	From   string
	To     string
	Amount int // In ¢
}
