package matrix

import (
	"testing"
)

func newMatrix(b *testing.B) (*Matrix, float64) {
	// Looking at production code, this is a representative matrix shape.
	m, err := New(10_000, 64)
	if err != nil {
		b.Fatal(err)
	}

	sum := 0.0
	for r := range m.Rows {
		for c := range m.Cols {
			v := float64((r + 1) * (c + 1))
			i := m.index(r, c)
			m.data[i] = v
			sum += v
		}
	}

	return m, sum
}

func BenchmarkSum(b *testing.B) {
	m, expected := newMatrix(b)
	b.ResetTimer()

	for b.Loop() {
		s := m.Sum()
		if expected != s {
			b.Fatal(s)
		}
	}
}
