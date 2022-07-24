package multithread_sort

import (
	"algoritmos-tcc-golang/constants"
	"sync"
)

type quickSort struct {
	threadsCount int
}

func NewMultithreadQuickSort() *quickSort {
	return &quickSort{
		threadsCount: 0,
	}
}

func (ss *quickSort) Sort(array []int) {
	ss.quickSort(array, 0, len(array)-1)
}

func (ss *quickSort) quickSort(array []int, low int, high int) {
	if low < high {
		pi := ss.partition(array, low, high)

		localThreadsCount := 0

		var wg sync.WaitGroup

		if ss.threadsCount < constants.THREADS_NUMBER {
			ss.threadsCount++
			localThreadsCount++
			wg.Add(1)

			go func() {
				ss.quickSort(array, low, pi-1)

				wg.Done()
			}()
		} else {
			ss.quickSort(array, low, pi-1)
		}

		if ss.threadsCount < constants.THREADS_NUMBER {
			ss.threadsCount++
			localThreadsCount++
			wg.Add(1)

			go func() {
				ss.quickSort(array, pi+1, high)

				wg.Done()
			}()
		} else {
			ss.quickSort(array, pi+1, high)
		}

		wg.Wait()
		ss.threadsCount -= localThreadsCount
	}
}

func (ss *quickSort) partition(array []int, low int, high int) int {
	pivot := array[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[high] = array[high], array[i+1]

	return i + 1
}
