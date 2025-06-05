package main

import (
	"encoding/json"
	"fmt"
)

// Zero vs Missing Value

func main() {
	/*
		m := map[string]int{
			"A": 47,
			"B": 36,
		}
		v := m["C"]
		fmt.Println(v)

		v, ok := m["C"]
		fmt.Println(v, ok)
		return
	*/
	/*
		//csvFile := "c:\to\new\reports\2025.csv"
		// `a` is a "raw string", \ is just a \
		csvFile := `c:\to\new\reports\2025.csv`
		fmt.Println(csvFile)
		return
	*/
	// data := []byte(`{"image": "ubuntu", "count": 3}`)
	data := []byte(`{"image": "ubuntu"}`)
	// Business requirement: If user didn't send count, use 1.
	// Otherwise count must be bigger than 0

	/*
		var r StartVM
		if err := json.Unmarshal(data, &r); err != nil {
			fmt.Println("ERROR:", err)
			return
		}
	*/

	/* Solution 2: Use map[string]any
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	var r StartVM
	c, ok := m["count"]
	if !ok {
		r.Count = 1
	} else {
		i, ok := c.(int)
		if !ok {
			fmt.Printf("ERROR: bad count: %#v\n", c)
		}
		r.Count = i
	}
	// Same for image ...
	*/

	// Solution 3: Use defaults
	r := StartVM{Count: 1}
	if err := json.Unmarshal(data, &r); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	if r.Count <= 0 {
		fmt.Println("ERROR: bad count - ", r.Count)
	}
	fmt.Println(r)
}

/* API endpoint with several JSON messages formats
- mapstructure
- json.RawMessage, good if messages are in {"type": "login", "payload": { ... }}
type Request struct {
	Type string
	Payload json.RawMessage
}
*/

type StartVM struct {
	Image string
	// Solution 1: Use pointers
	Count int
}
