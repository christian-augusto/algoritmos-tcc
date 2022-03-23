package main

import (
	"algoritmos-tcc-golang/constants"
	"algoritmos-tcc-golang/monothread_sort"
	"algoritmos-tcc-golang/multithread_sort"
	"algoritmos-tcc-golang/utils"
	"fmt"
	"time"
)

func main() {
	for size := constants.MIN; size <= constants.MAX; size++ {
		fmt.Printf("size: %v\n", size)
		sort(size)
		separator()
	}
}

func sort(size int64) {
	array1 := utils.GenerateRandomInt64Array(size)
	array2 := make([]int64, size)
	array3 := make([]int64, size)
	copy(array2, array1)
	copy(array3, array1)

	if constants.LOGS {
		if constants.TOTAL_LOGS {
			fmt.Println(array1)
			fmt.Println(array2)
			fmt.Println(array3)
		}
	}

	// selectionSort

	start1 := time.Now()

	selectionSort := monothread_sort.NewSelectionSort(array1)

	selectionSort.Sort()

	elapsed1 := time.Since(start1)

	// mergeSort

	start2 := time.Now()

	mergeSort := monothread_sort.NewMergeSort(array2)

	mergeSort.Sort()

	elapsed2 := time.Since(start2)

	// multithreadMergeSort

	start3 := time.Now()

	multithreadMergeSort := multithread_sort.NewMultithreadMergeSort()

	multithreadMergeSort.Sort(array3)

	elapsed3 := time.Since(start3)

	if constants.LOGS {
		fmt.Printf("Selection sort execution time: %v\n", elapsed1)
		fmt.Printf("Merge sort execution time: %v\n", elapsed2)
		fmt.Printf("Multithread merge sort execution time: %v\n", elapsed3)

		if constants.TOTAL_LOGS {
			fmt.Println(array1)
			fmt.Println(array2)
			fmt.Println(array3)
		}
	}
}

func separator() {
	fmt.Print("\n\n\n")
}
