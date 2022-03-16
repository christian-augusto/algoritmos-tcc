package main

import (
	"algoritmos-tcc-golang/monothread_sort"
	"algoritmos-tcc-golang/utils"
	"fmt"
	"time"
)

func main() {
	logs := true
	total_logs := false
	max := int64(10)
	// max := int64(100)
	// max := int64(1_000)
	// max := int64(10_000)
	// max := int64(100_000)
	// max := int64(1_000_000)

	for size := int64(1); size <= max; size++ {
		fmt.Printf("size: %v\n", size)
		sort(size, logs, total_logs)
		separator()
	}
}

func sort(size int64, logs bool, total_logs bool) {
	array1 := utils.GenerateRandomInt64Array(size)
	array2 := make([]int64, size)
	copy(array2, array1)

	if logs {
		if total_logs {
			fmt.Println(array1)
			fmt.Println(array2)
		}
	}

	start1 := time.Now()

	selectionSort := monothread_sort.NewSelectionSort(array1)

	selectionSort.Sort()

	elapsed1 := time.Since(start1)

	fmt.Printf("Selection sort execution time: %v\n", elapsed1)

	start2 := time.Now()

	mergeSort := monothread_sort.NewMergeSort(array2)

	mergeSort.Sort()

	elapsed2 := time.Since(start2)

	fmt.Printf("Merge sort execution time: %v\n", elapsed2)

	if logs {
		if total_logs {
			fmt.Println(array1)
			fmt.Println(array2)
		}
	}
}

func separator() {
	fmt.Print("\n\n\n")
}
