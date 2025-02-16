package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"shortest-path/internal"
)

func printMatrixWithPath(matrix [][]string, path []internal.Point) {
	pathSet := make(map[internal.Point]bool, len(path))
	for _, p := range path {
		pathSet[p] = true
	}

	for i, row := range matrix {
		for j, cell := range row {
			if pathSet[internal.NewPoint(i, j)] && cell != "E" {
				fmt.Print("* ")
			} else {
				fmt.Printf("%s ", cell)
			}
		}
		fmt.Println()
	}
}

func readMatrix(filename string) ([][]string, internal.Point, internal.Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, internal.Point{}, internal.Point{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := make([][]string, 0)
	var start, end internal.Point

	numRows := 0
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		matrix = append(matrix, row)

		for yPos, char := range row {
			if char == "S" {
				start = internal.NewPoint(numRows, yPos)
			} else if char == "E" {
				end = internal.NewPoint(numRows, yPos)
			}
		}
		numRows++
	}

	if err = scanner.Err(); err != nil {
		return nil, internal.Point{}, internal.Point{}, err
	}

	return matrix, start, end, nil
}

func main() {
	filename := "shortest-path/files/matrix.txt"
	matrix, start, end, err := readMatrix(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении матрицы:", err)
		return
	}

	if start == end {
		fmt.Println("Невалидная матрица: координаты начала равны координатам конца")
	}

	distance, path := internal.ParallelSearch(matrix, start, end)
	if distance == -1 {
		fmt.Println("Пути не найдено")
	} else {
		fmt.Printf("Длина кратчайшего пути: %d\n", distance)
		fmt.Println("Путь:")
		for _, p := range path[:len(path)-1] {
			fmt.Printf("%s -> ", p.ToString())
		}
		fmt.Printf("%s\n", path[len(path)-1].ToString())
		printMatrixWithPath(matrix, path)
	}
}
