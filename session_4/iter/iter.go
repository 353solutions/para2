package main

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"os"
	"runtime"
	"time"
)

func main() {
	for i := range 4 { // Generate values
		fmt.Println(i) // Use values
		if i > 1 {
			break
		}
	}

	yield := func(n int) bool {
		fmt.Println(n)
		return n <= 1
	}

	for i := range 4 {
		if !yield(i) {
			break
		}
	}

	n := 0
	for t := range Tick(time.Second) {
		fmt.Println("t:", t)
		n++
		if n > 3 {
			break
		}
	}

	file, err := os.Open("logs.json.gz")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	start := time.Now()
	/*
		logs, err := LoadLogsSlice(gz)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		count := len(logs)
	*/
	count := 0
	seq := LoadLogs(gz)
	seq = Filter(seq, IsValid)
	for range seq {
		count++
	}
	duration := time.Since(start)
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	alloc_mb := float64(mem.Alloc) / (1 << 20)
	fmt.Printf("%d logs in %v with %.2fMB\n", count, duration, alloc_mb)
}

/*
iter := func(v T) bool {
	// call user code with v
	return false // return true if there was a break/return in the "for loop"
}
*/

// Before filter: 5350200
// After filter:  5350195

func Filter[T any](in iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	fn := func(yield func(T) bool) {
		for v := range in {
			if !pred(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}

	return fn
}

func Head[T any](s iter.Seq[T], n int) iter.Seq[T] {
	// TODO: n > 0
	fn := func(yield func(T) bool) {
		for v := range s {
			if !yield(v) {
				return
			}
			n--
			if n == 0 {
				return
			}
		}
	}

	return fn
}

// Valid logs out of first 10 logs (BUG?)
// logs := loadLogs()
// logs = Head(logs, 10)
// logs = Filter(logs, IsValid)

// First 10 valid logs
// logs := loadLogs()
// logs = Filter(logs, IsValid)
// logs = Head(logs, 10)

/* Stream operations
- Map: [1, 2, 3], func(n int) int {return n * 2} -> [2 4 6]
- Filter: [1 2 3], func(n int) bool {return n < 3} -> [1 2]
- Reduce: [1 2 3], sum -> 6
*/

// Load and query logs

func LoadLogs(r io.Reader) iter.Seq[Log] {
	fn := func(yield func(Log) bool) {
		dec := json.NewDecoder(r)
		for {
			var l Log
			err := dec.Decode(&l)
			if errors.Is(err, io.EOF) {
				return
			}

			if err != nil {
				l = Log{}
			}

			if !yield(l) {
				return
			}
		}
	}

	return fn

}

func LoadLogsSlice(r io.Reader) ([]Log, error) {
	var logs []Log
	dec := json.NewDecoder(r)
	for {
		var l Log
		err := dec.Decode(&l)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}

	return logs, nil
}

type Log struct {
	Host    string    `json:"host"`
	Time    time.Time `json:"time"`
	Request string    `json:"request"`
	Status  int       `json:"status"`
	Bytes   int       `json:"bytes"`
}

func IsValid(log Log) bool {
	switch {
	case log.Host == "":
		return false
	case log.Time.IsZero():
		return false
	case log.Request == "":
		return false
	case log.Status < 100 || log.Status >= 600:
		return false
	case log.Bytes < 0:
		return false
	}

	return true
}

func Tick(d time.Duration) iter.Seq[time.Time] {
	fn := func(yield func(time.Time) bool) {
		for {
			if !yield(time.Now()) {
				return
			}

			time.Sleep(d)
		}
	}

	return fn
}
