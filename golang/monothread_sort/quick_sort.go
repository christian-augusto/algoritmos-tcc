package monothread_sort

type quickSort struct{}

func NewQuickSort() *quickSort {
	return &quickSort{}
}

func (ss *quickSort) Sort(array []int) {
	ss.quickSort(array, 0, len(array)-1)
}

func (ss *quickSort) quickSort(array []int, low int, high int) {
	if low < high {
		pi := ss.partition(array, low, high)

		ss.quickSort(array, low, pi-1)
		ss.quickSort(array, pi+1, high)
	}
}

func (ss *quickSort) partition(array []int, low int, high int) int {
	pivot := array[high]

	i := (low - 1)

	for j := low; j <= high-1; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[high] = array[high], array[i+1]

	return i + 1
}
