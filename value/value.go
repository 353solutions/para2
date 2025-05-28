package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
Types

JSON <-> Go
string <-> string
null <-> nil
number <-> float64, float32, int, int8 ... int64, uint8 ... uint64
true/false <-> bool
array <-> []T, []any
object <-> struct, map[string]any

MIA:
- time.Time -> string in RFC3339 format
- []byte -> string in base64
*/

func main() {
	n := 19
	data, err := json.Marshal(n)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	var a any
	if err := json.Unmarshal(data, &a); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(n == a)
	fmt.Printf("a: %v (%T)\n", a, a)

	var n2 int
	if err := json.Unmarshal(data, &n2); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(n == n2)

	if err := json.Unmarshal(data, &a); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))

	t := time.Now()
	data, err = json.Marshal(t)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))

	var t2 time.Time
	if err := json.Unmarshal(data, &t2); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("time:", t == t2)
	fmt.Println(t)
	fmt.Println(t2)
	fmt.Println("time:", t.Equal(t2))

	// Time in computers
	// wall clock: human
	// monotonic: machine, going up only

}
