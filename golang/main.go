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
	array_1 := utils.GenerateRandomIntArray(size, constants.RANDOM_INT_MIN, constants.RANDOM_INT_MAX)
	array_2 := make([]int, size)
	array_3 := make([]int, size)
	array_4 := make([]int, size)
	array_5 := make([]int, size)
	copy(array_2, array_1)
	copy(array_3, array_1)
	copy(array_4, array_1)
	copy(array_5, array_1)

	if constants.LOGS {
		if constants.TOTAL_LOGS {
			fmt.Println(array_1)
		}
	}

	// selection sort

	selectionSort := monothread_sort.NewSelectionSort()

	start_1 := time.Now().UnixMilli()

	selectionSort.Sort(array_1)

	elapsed_1 := time.Now().UnixMilli() - start_1

	// merge sort

	mergeSort := monothread_sort.NewMergeSort()

	start_2 := time.Now().UnixMilli()

	mergeSort.Sort(array_2)

	elapsed_2 := time.Now().UnixMilli() - start_2

	// quick sort

	quickSort := monothread_sort.NewQuickSort()

	start_3 := time.Now().UnixMilli()

	quickSort.Sort(array_3)

	elapsed_3 := time.Now().UnixMilli() - start_3

	// multithread merge sort

	multithreadMergeSort := multithread_sort.NewMultithreadMergeSort()

	start_4 := time.Now().UnixMilli()

	multithreadMergeSort.Sort(array_4)

	elapsed_4 := time.Now().UnixMilli() - start_4

	// multithread quick sort

	multithreadQuickSort := multithread_sort.NewMultithreadQuickSort()

	start_5 := time.Now().UnixMilli()

	multithreadQuickSort.Sort(array_5)

	elapsed_5 := time.Now().UnixMilli() - start_5

	if constants.LOGS {
		fmt.Printf("Selection sort execution time: %v ms\n", elapsed_1)
		fmt.Printf("Merge sort execution time: %v ms\n", elapsed_2)
		fmt.Printf("Quick sort execution time: %v ms\n", elapsed_3)
		fmt.Printf("Multithread merge sort execution time: %v ms\n", elapsed_4)
		fmt.Printf("Multithread quick sort execution time: %v ms\n", elapsed_5)

		if constants.TOTAL_LOGS {
			fmt.Println(array_1)
			fmt.Println(array_2)
			fmt.Println(array_3)
			fmt.Println(array_4)
			fmt.Println(array_5)
		}
	}
}

func separator() {
	fmt.Print("\n\n\n")
}
