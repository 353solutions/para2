package main

import (
	"fmt"
	"time"
)

func main() {
	signs := `-\|/`
	i := 0
	for range 20 {
		fmt.Printf("working %c\r", signs[i])
		time.Sleep(300 * time.Millisecond)
		i = (i + 1) % len(signs)
	}

}
