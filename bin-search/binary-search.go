package main

import "sync"

const NotFoundIdx = -1

func Search(sortedArr []int, target int) int {
	leftIdx, rightIdx := 0, len(sortedArr)-1

	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2

		if sortedArr[midIdx] == target {
			return midIdx
		}

		if sortedArr[midIdx] < target {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}
	}

	return -1
}

func ParallelSearch(sortedArr []int, targets []int) []int {
	result := make([]int, len(targets))
	wg := sync.WaitGroup{}
	for i, target := range targets {
		go func(i, target int) {
			defer wg.Done()

			wg.Add(1)
			result[i] = Search(sortedArr, target)
		}(i, target)
	}

	wg.Wait()
	return result
}
