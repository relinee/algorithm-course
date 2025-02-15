package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	Value [][]float64
}

func NewMatrix(matrix [][]float64) *Matrix {
	return &Matrix{
		Value: matrix,
	}
}

func CreateFromFile(filename string) (result *Matrix, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			result = nil
		}
	}(file)

	matrix := make([][]float64, 0)
	s := bufio.NewScanner(file)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		row := make([]float64, len(fields))
		matrix = append(matrix, row)
		for i, field := range fields {
			row[i], err = strconv.ParseFloat(field, 64)
			if err != nil {
				return nil, err
			}
		}
	}

	if err = s.Err(); err != nil {
		return nil, err
	}

	return NewMatrix(matrix), nil
}

func (m *Matrix) IsValidForGaussMethod() bool {
	return !m.isSquareMatrix()
}

func (m *Matrix) isSquareMatrix() bool {
	n := len(m.Value)
	for _, row := range m.Value {
		if len(row) != n+1 {
			return false
		}
	}
	return true
}
