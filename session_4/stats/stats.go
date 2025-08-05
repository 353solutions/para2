package main

import (
	"cmp"
	"fmt"
	"time"
)

func main() {
	fmt.Println(Max([]int{3, 1, 2}))     // call Max[int]
	fmt.Println(Max([]float64{3, 1, 2})) // call Max[float64]
	fmt.Println(Max[int](nil))
	fmt.Println(Max([]time.Month{time.March, time.January, time.February}))

	m, err := NewMatrix[float64](10, 3)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(m.At(3, 2))

	scores := map[string]int{
		"Rick":   10_000,
		"Morty":  90,
		"Summer": 8_900,
		"Beth":   11_322,
	}
	v, err := MaxMap(scores)
	fmt.Printf("max: %v (err=%v)\n", v, err) // Beth (err=nil)
}

/*
Exercise: Write MaxMap (generic function) that gets a map where values are
number and return the key the has the maximal corresponding value.

You can use cmp.Ordered for the values
*/
func MaxMap[K comparable, V cmp.Ordered](m map[K]V) (K, error) {
	if len(m) == 0 {
		return zero[K](), fmt.Errorf("max on empty map")
	}

	var (
		first = true
		maxK  K
		maxV  V
	)

	for k, v := range m {
		if first || v > maxV {
			maxK, maxV = k, v
		}
		first = false
	}

	return maxK, nil
}

/*
Go generics uses "gc shape stenciling"
- Two types have the same shape if the underlying type in memory has the same shape
- All pointers have the same gc shape

When to use generics:
- Function with identical code, different signature
- Generic containers
*/

func (m *Matrix[T]) At(row, col int) (T, error) {
	i := (row * m.Cols) + col
	if i > len(m.data) {
		return zero[T](), fmt.Errorf("%d/%d out of range for %d/%d", row, col, m.Rows, m.Cols)
	}

	return m.data[i], nil
}

type Number interface {
	~int | ~float64
}

func NewMatrix[T Number](rows, cols int) (*Matrix[T], error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("%d/%d - bad dimension", rows, cols)
	}

	m := Matrix[T]{
		Rows: rows,
		Cols: cols,
		data: make([]T, rows*cols),
	}

	return &m, nil
}

type Matrix[T Number] struct {
	Rows int
	Cols int

	data []T
}

type Ordered interface {
	~int | ~float64 | ~string
}

func zero[T any]() T {
	var v T
	return v
}

// T is a "type constraint", not a new type
func Max[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		//var zero T
		//return zero, fmt.Errorf("max of empty slice")
		return zero[T](), fmt.Errorf("max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

/*
func MaxInts(values []int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

func MaxFloat64(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

*/
