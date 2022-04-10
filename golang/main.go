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

func sort(size int) {
	array1 := utils.GenerateRandomIntArray(size)
	array2 := make([]int, size)
	array3 := make([]int, size)
	array4 := make([]int, size)
	copy(array2, array1)
	copy(array3, array1)
	copy(array4, array1)

	if constants.LOGS {
		if constants.TOTAL_LOGS {
			fmt.Println(array1)
		}
	}

	// selectionSort

	start1 := time.Now()

	selectionSort := monothread_sort.NewSelectionSort()

	selectionSort.Sort(array1)

	elapsed1 := time.Since(start1)

	// mergeSort

	start2 := time.Now()

	mergeSort := monothread_sort.NewMergeSort()

	mergeSort.Sort(array2)

	elapsed2 := time.Since(start2)

	// multithreadMergeSort

	start3 := time.Now()

	multithreadMergeSort := multithread_sort.NewMultithreadMergeSort()

	multithreadMergeSort.Sort(array3)

	elapsed3 := time.Since(start3)

	// quickSort

	start4 := time.Now()

	quickSort := monothread_sort.NewQuickSort()

	quickSort.Sort(array4)

	elapsed4 := time.Since(start4)

	if constants.LOGS {
		fmt.Printf("Selection sort execution time: %v\n", elapsed1)
		fmt.Printf("Merge sort execution time: %v\n", elapsed2)
		fmt.Printf("Multithread merge sort execution time: %v\n", elapsed3)
		fmt.Printf("Quick sort execution time: %v\n", elapsed4)

		if constants.TOTAL_LOGS {
			fmt.Println(array1)
			fmt.Println(array2)
			fmt.Println(array3)
			fmt.Println(array4)
		}
	}
}

func separator() {
	fmt.Print("\n\n\n")
}
