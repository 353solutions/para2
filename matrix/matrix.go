package matrix

import "fmt"

// Matrix is a 2-dimensional float64 matrix.
type Matrix struct {
	Rows int
	Cols int
	data []float64
}

// New returns a new Matrix with dimension (`rows`, `cols`).
func New(rows, cols int) (*Matrix, error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("bad dimension: %d/%d", rows, cols)
	}

	m := Matrix{
		Rows: rows,
		Cols: cols,
		data: make([]float64, rows*cols),
	}
	return &m, nil
}

func (m *Matrix) index(row, col int) int {
	return (row * m.Cols) + col
}

// Sum returns the sum of all the values in m.
func (m *Matrix) Sum() float64 {
	total := 0.0
	for col := range m.Cols {
		for row := range m.Rows {
			i := m.index(row, col)
			total += m.data[i]
		}
	}

	return total
}
