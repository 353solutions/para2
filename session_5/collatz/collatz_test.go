package collatz

import (
	"fmt"
	"testing"
)

var collatzCases = []struct {
	n int
	c int
}{
	{13, 40},
	{40, 20},
	{20, 10},
	{10, 5},
	{5, 16},
	{16, 8},
}

func TestCollatz(t *testing.T) {
	for _, tc := range collatzCases {
		name := fmt.Sprintf("%d", tc.n)
		t.Run(name, func(t *testing.T) {
			if c := Collatz(tc.n); c != tc.c {
				t.Fatalf("%d: wanted: %d, got: %d", tc.n, tc.c, c)
			}
		})
	}
}

func TestCollatzLen(t *testing.T) {
	n := CollatzLen(13)
	const want = 10
	if n != want {
		t.Fatalf("want: %d, got %d", want, n)
	}
}

func TestMaxCollatz(t *testing.T) {
	n, _ := MaxCollatz(1000)
	const want = 871
	if n != want {
		t.Fatalf("want: %d, got: %d", want, n)
	}
}
