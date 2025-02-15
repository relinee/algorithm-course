package main

import (
	"fmt"

	"parallel-gauss/internal"
)

func main() {
	matrix, err := internal.CreateFromFile("parallel-gauss/files/matrix.txt")
	if err != nil {
		panic(fmt.Errorf("ошибка при чтении матрицы: %v", err))
	}

	solution, err := matrix.CalculateSolveByGaussElimination()
	if err != nil {
		panic(fmt.Errorf("ошибка при поиске решения: %v", err))
	}

	fmt.Println("Решение системы:")
	for i, val := range solution {
		fmt.Printf("x%d = %f\n", i+1, val)
	}
}
