package main

import "fmt"

func main() {
	sortedArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := []int{4, 5}

	result := ParallelSearch(sortedArr, target)

	for i := range result {
		if result[i] == NotFoundIdx {
			fmt.Println(target[i], "не найден")
			continue
		}
		fmt.Println(target[i], "имеет позицию", result[i])
		fmt.Println(result[i], result[i])
	}
}
