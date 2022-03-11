package main

import (
	"algoritmos-tcc-golang/monothread_sort"
	"algoritmos-tcc-golang/utils"
	"fmt"
	"time"
)

func main() {
	size := int64(10)
	// size := int64(1_000)
	// size := int64(100_000)
	// size := int64(1_000_000)

	array := utils.GenerateRandomInt64Array(size)

	// fmt.Println(array)
	fmt.Println(array[0])

	start := time.Now()

	selectionSort := monothread_sort.NewSelectionSort(array)

	selectionSort.Sort()

	elapsed := time.Since(start)

	fmt.Printf("Execution time: %v\n", elapsed)

	// fmt.Println(array)
	fmt.Println(array[size-1])
}
