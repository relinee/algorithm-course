package internal

import (
	"fmt"
	"math"
	"sync"
)

func (m *Matrix) CalculateSolveByGaussElimination() ([]float64, error) {
	if m.IsValidForGaussMethod() {
		return nil, fmt.Errorf("матрица не подходит для использования метода Гаусса")
	}

	n := len(m.Value)

	// Прямой ход
	for i := 0; i < n; i++ {
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(m.Value[k][i]) > math.Abs(m.Value[maxRow][i]) {
				maxRow = k
			}
		}

		if m.Value[maxRow][i] == 0 {
			return nil, fmt.Errorf("матрица вырождена, система не имеет единственного решения")
		}

		m.Value[i], m.Value[maxRow] = m.Value[maxRow], m.Value[i]

		var wg sync.WaitGroup
		for k := i + 1; k < n; k++ { // по строкам
			wg.Add(1)
			go func(k int) {
				defer wg.Done()
				factor := m.Value[k][i] / m.Value[i][i]
				for j := i; j < n+1; j++ { // по столбцам
					m.Value[k][j] -= factor * m.Value[i][j]
				}
			}(k)
		}
		wg.Wait()
	}

	// Обратный ход
	result := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		result[i] = m.Value[i][n] / m.Value[i][i]
		for j := i - 1; j >= 0; j-- {
			m.Value[j][n] -= m.Value[j][i] * result[i]
		}
	}

	return result, nil
}
