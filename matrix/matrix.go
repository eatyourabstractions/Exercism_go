package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(ms string) (*Matrix, error) {
	rows := strings.Split(ms, "\n")
	matrix := make(Matrix, len(rows))
	var valueLen int
	for r, row := range rows {
		values := strings.Split(strings.TrimSpace(row), " ")
		switch {
		case r == 0:
			valueLen = len(values)
		case valueLen != len(values):
			return nil, errors.New("uneven rows")
		}
		matrix[r] = make([]int, len(values))
		for c, v := range values {
			var err error
			matrix[r][c], err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

		}
	}
	return &matrix, nil
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(*m))
	for r, row := range *m {
		rows[r] = make([]int, len(row))
		copy(rows[r], row)
	}
	return rows
}

func (m *Matrix) Cols() [][]int {
	cols := make([][]int, len((*m)[0]))
	for c := range cols {
		cols[c] = make([]int, len(*m))
		for r := range cols[c] {
			cols[c][r] = (*m)[r][c]
		}
	}
	return cols
}

func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(*m) || c >= len((*m)[0]) {
		return false
	}
	(*m)[r][c] = val
	return true
}