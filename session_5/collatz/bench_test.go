package collatz

import "testing"

func BenchmarkMaxCollatz(b *testing.B) {
	const want = 871
	for b.Loop() {
		n, _ := MaxCollatz(1000)
		if n != want {
			b.Fatalf("want: %d, got: %d", want, n)
		}
	}
}
