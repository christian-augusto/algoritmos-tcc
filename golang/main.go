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
	for _, size := range constants.ARRAY_SIZES {
		fmt.Printf("size: %v\n", size)
		sort(size)
		separator()
	}
}

func sort(size int) {
	array1 := utils.GenerateRandomIntArray(size, constants.RANDOM_INT_MIN, constants.RANDOM_INT_MAX)
	array2 := make([]int, size)
	array3 := make([]int, size)
	array4 := make([]int, size)
	array5 := make([]int, size)
	copy(array2, array1)
	copy(array3, array1)
	copy(array4, array1)
	copy(array5, array1)

	if constants.LOGS {
		if constants.TOTAL_LOGS {
			fmt.Println(array1)
		}
	}

	// selection sort

	start1 := time.Now().UnixMilli()

	selectionSort := monothread_sort.NewSelectionSort()

	selectionSort.Sort(array1)

	elapsed1 := time.Now().UnixMilli() - start1

	// merge sort

	start2 := time.Now().UnixMilli()

	mergeSort := monothread_sort.NewMergeSort()

	mergeSort.Sort(array2)

	elapsed2 := time.Now().UnixMilli() - start2

	// quick sort

	start3 := time.Now().UnixMilli()

	quickSort := monothread_sort.NewQuickSort()

	quickSort.Sort(array3)

	elapsed3 := time.Now().UnixMilli() - start3

	// multithread merge sort

	start4 := time.Now().UnixMilli()

	multithreadMergeSort := multithread_sort.NewMultithreadMergeSort()

	multithreadMergeSort.Sort(array4)

	elapsed4 := time.Now().UnixMilli() - start4

	// multithread quick sort

	start5 := time.Now().UnixMilli()

	multithreadQuickSort := multithread_sort.NewMultithreadQuickSort()

	multithreadQuickSort.Sort(array5)

	elapsed5 := time.Now().UnixMilli() - start5

	if constants.LOGS {
		fmt.Printf("Selection sort execution time: %v ms\n", elapsed1)
		fmt.Printf("Merge sort execution time: %v ms\n", elapsed2)
		fmt.Printf("Quick sort execution time: %v ms\n", elapsed3)
		fmt.Printf("Multithread merge sort execution time: %v ms\n", elapsed4)
		fmt.Printf("Multithread quick sort execution time: %v ms\n", elapsed5)

		if constants.TOTAL_LOGS {
			fmt.Println(array1)
			fmt.Println(array2)
			fmt.Println(array3)
			fmt.Println(array4)
			fmt.Println(array5)
		}
	}
}

func separator() {
	fmt.Print("\n\n\n")
}
