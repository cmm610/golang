package matrix

import (
	"strings"
	"strconv"
	"errors"
)

type matrix [][]int

func New(s string) (*matrix, error) {
	var m matrix
	for _, line := range strings.Split(s, "\n") {
		fields := strings.Fields(line)
		if len(m) > 0 && len(fields) != len(m[len(m)-1]) {
			return nil, errors.New("wrong number of fields")
		}
		row := make([]int, len(fields))
		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, errors.New("not a number")
			}
			row[i] = num
		}
		m = append(m, row)
	}
	return &m, nil
}

func (m *matrix) Cols() [][]int {
	var nCols int
	if len(*m) > 0 {
		nCols = len((*m)[0])
	}
	n := make([][]int, nCols)
	for _, row := range *m {
		for i, col := range row {
			n[i] = append(n[i], col)
		}
	}
	return n
}

func (m *matrix) Rows() [][]int {
	n := make([][]int, len(*m))
	for i, row := range *m {
		n[i] = make([]int, len(row))
		for j, val := range row {
			n[i][j] = val
		}
	}
	return n
}

func (m *matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(*m) || (len(*m) > 0 && c >= len((*m)[0])) {
		return false
	}
	(*m)[r][c] = val
	return true
}