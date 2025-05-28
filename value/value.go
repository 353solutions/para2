package main

import (
	"bytes"
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

encoding/json API
JSON -> []byte -> Go: json.Unmarshal
Go -> []byte -> JSON: json.Marshal
JSON -> io.Reader -> Go: json.Decoder
Go -> io.Writer -> JSON: json.Encoder

You can use bytes.Buffer or bytes.Reader for in memory io.Reader/Writer
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

	v := Value{
		Unit:   "cm",
		Amount: 136.7,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&v)
	fmt.Println(buf.String())
}

// What json.Marshal does
func Marshal(val any) ([]byte, error) {
	if m, ok := val.(json.Marshaler); ok {
		return m.MarshalJSON()
	}

	// do the defaults
	return nil, nil

}

// Business: Show Value as JSON string in 136.7cm format

/* In order for encoding/json to use MarshalJSON, the receiver type must match
method: value, encoded: value: OK
method: value, encode: pointer: OK
method: pointer, encode: pointer: OK
method: pointer, encode: value: NO
*/

// MarshalJSON implements json.Marshaler
func (v Value) MarshalJSON() ([]byte, error) {
	// Step 1: Convert to a type encoding/json can handle
	s := fmt.Sprintf("%f%s", v.Amount, v.Unit)

	// Step 2: Use json.Marshal
	return json.Marshal(s)

	// Step 3: There is no step 3
}

type Value struct {
	Unit   string  `json:"unit,omitempty"` // field tag
	Amount float64 `json:"amount,omitempty"`
}

/*
API			stable: type User struct {}
Business    fast  : type User struct {}
Data		medium: type User struct {}

DRY: Do not repeat yourself
A little copying is better than a little dependency. (Go Team)
*/
